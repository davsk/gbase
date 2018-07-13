// /////////////////////////////////////////////////////////////
// 'structs.go'                                                 /
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

package config

// Owner base config interface for clients.
type Owner struct {
	Id    uint64 // 'toml:"OwnerId"'
	email string
}

// Ports base config interface for server io.
type Ports struct {
	Http  uint16
	Https uint16
	Rpc   uint16
}

// Server base config interface to find server.
type Server struct {
	Host string
	Ports
}

// Connect base config interface to connect.
type Connect struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

// Client config interface for all clients.
type Client struct {
	Title string
	Owner
	Game Server
}

// GameServer config interface for LAN  server.
type GameServer struct {
	Title string
	Ports
	Acct Server
	Game Connect
}

// AcctServer config interface for WAN server.
type AcctServer struct {
	Title string
	Ports
	Acct Connect
}

// TestServer config interface for test server.
type TestServer struct {
	Title      string
	GameServer Ports
	AcctServer Ports
	Game       Connect
	Acct       Connect
}
