// /////////////////////////////////////////////////////////////
// 'config.go'                                                 /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 12, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// package config loads config data from file
// or creates the file if it does not exist.
package config

import "davsk.net/gbase/pkg/toml_config"

// 'config_con_client.go'

const (
	// const RequiredTitle is base of filename
	// and Title of Toml file.
	aRequiredTitle = "config_con_client"
)

// ConfigConClient contains all config data to start program.
type ConfigConClient struct {
	toml_config.Client
}

func NewConfigConClient() ConfigConClient {
	var config ConfigConClient

	return config
}

func (*ConfigConClient) MustUpdate() {

}

// 'config_web_client.go'

const (
	// const RequiredTitle is base of filename
	// and Title of Toml file.
	bRequiredTitle = "config_web_client"
)

// ConfigWebClient contains all config data to start program.
type ConfigWebClient struct {
	Title string
}

func NewConfigWebClient() ConfigWebClient {
	var config ConfigWebClient

	return config
}

func (*ConfigWebClient) MustUpdate() {

}

// 'config_gui_client.go'

const (
	// const RequiredTitle is base of filename
	// and Title of Toml file.
	cRequiredTitle = "config_gui_client"
)

// ConfigGuiClient contains all config data to start program.
type ConfigGuiClient struct {
	Title string
}

func NewConfigGuiClient() ConfigGuiClient {
	var config ConfigGuiClient

	return config
}

func (*ConfigGuiClient) MustUpdate() {

}

// 'config_test.go'

const (
	// const RequiredTitle is base of filename
	// and Title of Toml file.
	dRequiredTitle = "config_test"
)

// ConfigTest contains all config data to start program.
type ConfigTest struct {
	Title string
}

func NewConfigTest() ConfigTest {
	var config ConfigTest

	return config
}

func (*ConfigTest) MustUpdate() {

}
