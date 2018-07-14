// /////////////////////////////////////////////////////////////
// 'GameServer.go'                                             /
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

import "davsk.net/gbase/pkg/toml_config"

const (
	// const kGsTitle is base of filename
	// and Title of Toml file.
	kGsTitle = "config_game_server"
)

// GameServer config interface for LAN  server.
type GameServer struct {
	Title string
	Ports
	Acct Server
	Game Connect
}

// NewGameServer creates GameServer with saved or default values.
func NewGameServer() GameServer {
	var gs GameServer

	// Load config from file.
	if err := toml_config.Load(kGsTitle, &gs); err != nil {
		// Save default config.
		gs.Default()
		gs.MustUpdate()
	}

	return ts
}

// Default GameServer
func (gs *GameServer) Default() {
	gs.Title = kGsTitle
	gs.Ports.Default()
	gs.Acct.Default("acct")
	gs.Game.Default("game")
}

// MustUpdate saves config, panics on fail.
func (gs *TestServer) MustUpdate() {
	toml_config.MustSave(kGsTitle, gs)
}
