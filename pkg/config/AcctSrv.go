// /////////////////////////////////////////////////////////////
// 'acctsrv.go'                                                /
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
	// const kAsTitle is base of filename
	// and Title of Toml file.
	kAsTitle = "config_acct_srv"
)

// AcctSrv is the config interface for WAN server.
type AcctSrv struct {
	Title string
	Ports
	Acct Connect
}

// NewAcctSrv creates acctsrv with saved or default values.
func NewAcctSrv() AcctSrv {
	var as AcctSrv

	// Load config from file.
	if err := tomlcfg.Load(kTsTitle, &as); err != nil {
		// Save default config.
		as.Default()
		as.MustUpdate()
	}

	return as
}

// Default AcctSrv receives title string.
func (as *AcctSrv) Default() {
	as.Title = kAsTitle
	as.Ports.Default()
	as.Acct.Default("acct")
}

// MustUpdate saves config, panics on fail.
func (as *AcctSrv) MustUpdate() {
	tomlcfg.MustSave(kAsTitle, as)
}
