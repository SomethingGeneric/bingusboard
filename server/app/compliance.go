package app

import "github.com/mattermost/focalboard/server/model"

// GetBoardsForCompliance retrieves boards for compliance reporting.
func (a *App) GetBoardsForCompliance(opts model.QueryBoardsForComplianceOptions) ([]*model.Board, bool, error) {
	return a.store.GetBoardsForCompliance(opts)
}

// GetBoardsComplianceHistory retrieves board history for compliance reporting.
func (a *App) GetBoardsComplianceHistory(opts model.QueryBoardsComplianceHistoryOptions) ([]*model.BoardHistory, bool, error) {
	return a.store.GetBoardsComplianceHistory(opts)
}

// GetBlocksComplianceHistory retrieves block history for compliance reporting.
func (a *App) GetBlocksComplianceHistory(opts model.QueryBlocksComplianceHistoryOptions) ([]*model.BlockHistory, bool, error) {
	return a.store.GetBlocksComplianceHistory(opts)
}
