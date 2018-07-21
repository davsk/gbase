// /////////////////////////////////////////////////////////////
// 'conclient.go'                                              /
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
	// const kCcTitle is base of filename
	// and Title of Toml file.
	kCcTitle = "config_con_client"
)

// ConClient contains all config data to start program.
type ConClient struct {
	Title string
	Owner
	Game Server
}

// NewConClient creates ConClient with saved or default values.
func NewConClient() ConClient {
	var cc ConClient

	// Load config from file.
	if err := tomlcfg.Load(kCcTitle, &cc); err != nil {
		// Save default config.
		cc.Default()
		cc.MustUpdate()
	}

	return cc
}

// Default ConClient.
func (cc *ConClient) Default() {
	cc.Title = kCcTitle
	cc.Owner.Default()
	cc.Game.Default("game")
}

// MustUpdate saves config, panics on fail.
func (cc *ConClient) MustUpdate() {
	tomlcfg.MustSave(kCcTitle, cc)
}
