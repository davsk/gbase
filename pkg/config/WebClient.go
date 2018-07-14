// /////////////////////////////////////////////////////////////
// 'WebClient.go'                                              /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on 7 14, 2018                                               /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

package config

import "davsk.net/gbase/pkg/toml_config"

const (
	// const kWcTitle is base of filename
	// and Title of Toml file.
	kWcTitle = "config_web_client"
)

// ConClient contains all config data to start program.
type WebClient struct {
	Title string
	Owner
	Game Server
}

// NewWebClient creates WebClient with saved or default values.
func NewWebClient() WebClient {
	var wc WebClient

	// Load config from file.
	if err := toml_config.Load(kWcTitle, &wc); err != nil {
		// Save default config.
		wc.Default()
		wc.MustUpdate()
	}

	return wc
}

// Default WebClient.
func (wc *WebClient) Default() {
	wc.Title = kWcTitle
	wc.Owner.Default()
	wc.Game.Default("game")
}

// MustUpdate saves config, panics on fail.
func (wc *WebClient) MustUpdate() {
	toml_config.MustSave(kWcTitle, wc)
}
