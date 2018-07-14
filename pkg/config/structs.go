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
