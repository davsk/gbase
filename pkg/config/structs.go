// /////////////////////////////////////////////////////////////
// 'structs.go'                                                /
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

// Default Owner.
func (o *Owner) Default() {
	o.Id = 0
	o.email = "email"
}

// Ports base config interface for server io.
type Ports struct {
	Http  uint16
	Https uint16
	Rpc   uint16
}

// Default Ports.
func (p *Ports) Default() {
	p.Http = 80
	p.Https = 433
	p.Rpc = 5000
}

// Server base config interface to find server.
type Server struct {
	Host string
	Ports
}

// Default Server receives choice string for host.
func (s *Server) Default(choice string) {
	// Default Host names
	switch choice {
	case "game":
		s.Host = "davsk.sytes.net"
		break
	case "acct":
		s.Host = "universe.gameserve.com"
		break
	default:
		s.Host = choice
	}

	// Default Ports
	s.Ports.Default()
}

// Connect base config interface to connect.
type Connect struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

// Default Connect receives choice of databases.
func (c *Connect) Default(choice string) {
	switch choice {
	case "game":
		c.Host = "localhost"
		c.Port = 5432
		c.Database = "universe"
		c.User = "postgres"
		c.Password = "password"
		break
	case "acct":
		c.Host = "localhost"
		c.Port = 5432
		c.Database = "business"
		c.User = "postgres"
		c.Password = "password"
		break
	default:
		c.Host = "localhost"
		c.Port = 5432
		c.Database = "postgres"
		c.User = "postgres"
		c.Password = "password"
	}
}

// Client config interface for all clients.
type Client struct {
	Title string
	Owner
	Game Server
}

// Default Client receives title.
func (c *Client) Default(title string) {
	c.Title = title
	c.Owner.Default()
	c.Game.Default("game")
}

// GameServer config interface for LAN  server.
type GameServer struct {
	Title string
	Ports
	Acct Server
	Game Connect
}

// Default GameServer
func (gs *GameServer) Default(title string) {
	gs.Title = title
	gs.Ports.Default()
	gs.Acct.Default("acct")
	gs.Game.Default("game")
}

// AcctServer config interface for WAN server.
type AcctServer struct {
	Title string
	Ports
	Acct Connect
}

// Default AcctServer receives title string.
func (as *AcctServer) Default(title string) {
	as.Title = title
	as.Ports.Default()
	as.Acct.Default("acct")
}
