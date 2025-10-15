package app

import (
	"io"
	"sync"
	"time"

	"github.com/mattermost/focalboard/server/auth"
	"github.com/mattermost/focalboard/server/services/config"
	"github.com/mattermost/focalboard/server/services/metrics"
	"github.com/mattermost/focalboard/server/services/notify"
	"github.com/mattermost/focalboard/server/services/permissions"
	"github.com/mattermost/focalboard/server/services/store"
	"github.com/mattermost/focalboard/server/services/webhook"
	"github.com/mattermost/focalboard/server/utils"
	"github.com/mattermost/focalboard/server/ws"

	mm_model "github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/mlog"
	"github.com/mattermost/mattermost/server/v8/platform/shared/filestore"
)

const (
	blockChangeNotifierQueueSize       = 1000
	blockChangeNotifierPoolSize        = 10
	blockChangeNotifierShutdownTimeout = time.Second * 10
)

type servicesAPI interface {
	GetUsersFromProfiles(options *mm_model.UserGetOptions) ([]*mm_model.User, error)
}

// ReadCloseSeeker is an interface for reading and seeking with close capability.
type ReadCloseSeeker = filestore.ReadCloseSeeker

type fileBackend interface {
	Reader(path string) (ReadCloseSeeker, error)
	FileExists(path string) (bool, error)
	CopyFile(oldPath, newPath string) error
	MoveFile(oldPath, newPath string) error
	WriteFile(fr io.Reader, path string) (int64, error)
	RemoveFile(path string) error
}

// Services provides the dependencies required by the App.
type Services struct {
	Auth             *auth.Auth
	Store            store.Store
	FilesBackend     fileBackend
	Webhook          *webhook.Client
	Metrics          *metrics.Metrics
	Notifications    *notify.Service
	Logger           mlog.LoggerIFace
	Permissions      permissions.PermissionsService
	SkipTemplateInit bool
	ServicesAPI      servicesAPI
}

// App represents the application layer of Focalboard.
type App struct {
	config              *config.Configuration
	store               store.Store
	auth                *auth.Auth
	wsAdapter           ws.Adapter
	filesBackend        fileBackend
	webhook             *webhook.Client
	metrics             *metrics.Metrics
	notifications       *notify.Service
	logger              mlog.LoggerIFace
	permissions         permissions.PermissionsService
	blockChangeNotifier *utils.CallbackQueue
	servicesAPI         servicesAPI

	cardLimitMux sync.RWMutex
	cardLimit    int
}

// SetConfig updates the application configuration.
func (a *App) SetConfig(config *config.Configuration) {
	a.config = config
}

// GetConfig returns the current application configuration.
func (a *App) GetConfig() *config.Configuration {
	return a.config
}

// New creates a new App instance with the provided configuration and services.
func New(config *config.Configuration, wsAdapter ws.Adapter, services Services) *App {
	app := &App{
		config:              config,
		store:               services.Store,
		auth:                services.Auth,
		wsAdapter:           wsAdapter,
		filesBackend:        services.FilesBackend,
		webhook:             services.Webhook,
		metrics:             services.Metrics,
		notifications:       services.Notifications,
		logger:              services.Logger,
		permissions:         services.Permissions,
		blockChangeNotifier: utils.NewCallbackQueue("blockChangeNotifier", blockChangeNotifierQueueSize, blockChangeNotifierPoolSize, services.Logger),
		servicesAPI:         services.ServicesAPI,
	}
	app.initialize(services.SkipTemplateInit)
	return app
}

// CardLimit returns the current card limit for the application.
func (a *App) CardLimit() int {
	a.cardLimitMux.RLock()
	defer a.cardLimitMux.RUnlock()
	return a.cardLimit
}

// SetCardLimit updates the card limit for the application.
func (a *App) SetCardLimit(cardLimit int) {
	a.cardLimitMux.Lock()
	defer a.cardLimitMux.Unlock()
	a.cardLimit = cardLimit
}

func (a *App) GetLicense() *mm_model.License {
	return a.store.GetLicense()
}
