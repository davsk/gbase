// /////////////////////////////////////////////////////////////
// 'TestServer.go'                                            /
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
	// const kTsTitle is base of filename
	// and Title of Toml file.
	kTsTitle = "config_test_server"
)

// TestServer config interface for test server.
type TestServer struct {
	Title      string
	GameServer Ports
	AcctServer Ports
	Game       Connect
	Acct       Connect
}

// NewTestServer creates TestServer with saved or default values.
func NewTestServer() TestServer {
	var ts TestServer

	// Load config from file.
	if err := toml_config.Load(kTsTitle, &ts); err != nil {
		// Save default config.
		ts.Default()
		ts.MustUpdate()
	}

	return ts
}

// Default TestServer receives title string.
func (ts *TestServer) Default() {
	ts.Title = kTsTitle
	ts.GameServer.Default()
	ts.AcctServer.Http = 8000
	ts.AcctServer.Https = 8080
	ts.AcctServer.Rpc = 5001
	ts.Game.Default("game")
	ts.Acct.Default("acct")
}

// MustUpdate saves config, panics on fail.
func (ts *TestServer) MustUpdate() {
	toml_config.MustSave(kTsTitle, ts)
}
