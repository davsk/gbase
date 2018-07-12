////////////////////////////////////////////////////////////////
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
////////////////////////////////////////////////////////////////

package toml_config

type Owner struct {
	Id    uint64 // 'toml:"OwnerId"'
	email string
}

type Ports struct {
	Http  uint16
	Https uint16
	Rpc   uint16
}

type Server struct {
	Host string
	Ports
}

type Client struct {
	Title string
	Owner
	GameServer Server
}

type GameServer struct {
	Title string
	Server
}

type AcctServer struct {
	Title string
	Server
}

type GameServices struct {
}

type AcctServices struct {
}
