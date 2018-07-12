// Copyright (c) 2018 Davsk℠. All Rights Reserved.
// Use of this source code is governed by an ISC License (ISC)
// that can be found in the LICENSE file.

// 'config.go'

// by David Skinner
// on July 11, 2018
// for Davsk℠ Universe 4.0 project gbase

// package config load config data from file or creates the file if it does not exist.
package config

// 'config_con_client.go'

// constants.
const (
	aRequiredTitle = "config_con_client"
)

// ConfigConClient contains all config data to start program.
type ConfigConClient struct {
	Title string
}

func NewConfigConClient() ConfigConClient {
	var config ConfigConClient

	return config
}

func (*ConfigConClient) MustUpdate() {

}

// 'config_web_client.go'

// constants.
const (
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

// constants.
const (
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

// constants.
const (
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

// 'config_game_server.go'

// constants.
const (
	eRequiredTitle = "config_game_server"
)

// ConfigGameServer contains all config data to start program.
type ConfigGameServer struct {
	Title string
}

func NewConfigGameServer() ConfigGameServer {
	var config ConfigGameServer

	return config
}

func (*ConfigGameServer) MustUpdate() {

}

// 'config_acct_server.go'

// constants.
const (
	fRequiredTitle = "config_acct_server"
)

// ConfigAcctServer contains all config data to start program.
type ConfigAcctServer struct {
	Title string
}

func NewConfigAcctServer() ConfigAcctServer {
	var config ConfigAcctServer

	return config
}

func (*ConfigAcctServer) MustUpdate() {

}

// 'config_test_server.go'

// constants.
const (
	gRequiredTitle = "config_test_server"
)

// ConfigTestServer contains all config data to start program.
type ConfigTestServer struct {
	Title string
}

func NewConfigTestServer() ConfigTestServer {
	var config ConfigTestServer

	return config
}

func (*ConfigTestServer) MustUpdate() {

}
