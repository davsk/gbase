// /////////////////////////////////////////////////////////////
// 'guiclient.go'                                              /
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
	// const kGcTitle is base of filename
	// and Title of Toml file.
	kGcTitle = "config_gui_client"
)

// guiclient contains all config data to start program.
type GuiClient struct {
	Title string
	Owner
	Game Server
}

// NewGuiClient creates guiclient with saved or default values.
func NewGuiClient() GuiClient {
	var gc GuiClient

	// Load config from file.
	if err := tomlcfg.Load(kGcTitle, &gc); err != nil {
		// Save default config.
		gc.Default()
		gc.MustUpdate()
	}

	return gc
}

// Default guiclient.
func (gc *GuiClient) Default() {
	gc.Title = kGcTitle
	gc.Owner.Default()
	gc.Game.Default("game")
}

// MustUpdate saves config, panics on fail.
func (gc *GuiClient) MustUpdate() {
	tomlcfg.MustSave(kGcTitle, gc)
}
