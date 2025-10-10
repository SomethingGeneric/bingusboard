// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

// GetUsedCardsCount returns the number of cards currently in use.
func (a *App) GetUsedCardsCount() (int, error) {
	return a.store.GetUsedCardsCount()
}
