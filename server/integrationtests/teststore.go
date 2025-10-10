package integrationtests

import (
	"github.com/mattermost/focalboard/server/services/store"

	mmModel "github.com/mattermost/mattermost/server/public/model"
)

type TestStore struct {
	store.Store
	license *mmModel.License
}

func newTestStoreWithLicense(store store.Store, isEnterprise bool) *TestStore {
	usersValue := 10000
	trueValue := true
	falseValue := false

	// Enterprise features that differ from Professional
	ldapEnabled := isEnterprise
	complianceEnabled := isEnterprise
	clusterEnabled := isEnterprise
	enterprisePluginsEnabled := isEnterprise
	remoteClusterEnabled := isEnterprise

	license := &mmModel.License{
		Features: &mmModel.Features{
			Users:                     &usersValue,
			LDAP:                      &ldapEnabled,
			LDAPGroups:                &ldapEnabled,
			MFA:                       &trueValue,
			GoogleOAuth:               &trueValue,
			Office365OAuth:            &trueValue,
			OpenId:                    &trueValue,
			Compliance:                &complianceEnabled,
			Cluster:                   &clusterEnabled,
			Metrics:                   &trueValue,
			MHPNS:                     &trueValue,
			SAML:                      &trueValue,
			Elasticsearch:             &trueValue,
			Announcement:              &trueValue,
			ThemeManagement:           &trueValue,
			EmailNotificationContents: &trueValue,
			DataRetention:             &trueValue,
			MessageExport:             &trueValue,
			CustomPermissionsSchemes:  &trueValue,
			CustomTermsOfService:      &trueValue,
			GuestAccounts:             &trueValue,
			GuestAccountsPermissions:  &trueValue,
			IDLoadedPushNotifications: &trueValue,
			LockTeammateNameDisplay:   &trueValue,
			EnterprisePlugins:         &enterprisePluginsEnabled,
			AdvancedLogging:           &trueValue,
			Cloud:                     &falseValue,
			SharedChannels:            &trueValue,
			RemoteClusterService:      &remoteClusterEnabled,
			FutureFeatures:            &trueValue,
		},
	}

	return &TestStore{
		Store:   store,
		license: license,
	}
}

func NewTestEnterpriseStore(store store.Store) *TestStore {
	return newTestStoreWithLicense(store, true)
}

func NewTestProfessionalStore(store store.Store) *TestStore {
	return newTestStoreWithLicense(store, false)
}

func (s *TestStore) GetLicense() *mmModel.License {
	return s.license
}
