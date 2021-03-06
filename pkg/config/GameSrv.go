// /////////////////////////////////////////////////////////////
// 'gameserver.go'                                             /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 14, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

package config

import "davsk.net/gbase/pkg/tomlcfg"

const (
	// const kGsTitle is base of filename
	// and Title of Toml file.
	kGsTitle = "config_game_server"
)

// gameserver config interface for LAN  server.
type GameSrv struct {
	Title string
	Service
	Acct Server
	Game Connect
}

// NewGameSrv creates gameserver with saved or default values.
func NewGameSrv() GameSrv {
	var gs GameSrv

	// Load config from file.
	if err := tomlcfg.Load(kGsTitle, &gs); err != nil {
		// Save default config.
		gs.Default()
		gs.MustUpdate()
	}

	return gs
}

// Default gameserver
func (gs *GameSrv) Default() {
	gs.Title = kGsTitle
	gs.Service.Default("universe.servegame.com")
	gs.Acct.Default("acct")
	gs.Game.Default("game")
}

// MustUpdate saves config, panics on fail.
func (gs *GameSrv) MustUpdate() {
	tomlcfg.MustSave(kGsTitle, gs)
}
