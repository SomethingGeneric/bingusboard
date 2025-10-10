package app

import (
	mm_model "github.com/mattermost/mattermost/server/public/model"
)

// HasPermissionToBoard checks if a user has a specific permission for a board.
func (a *App) HasPermissionToBoard(userID, boardID string, permission *mm_model.Permission) bool {
	return a.permissions.HasPermissionToBoard(userID, boardID, permission)
}
