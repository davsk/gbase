// /////////////////////////////////////////////////////////////
// 'Server.go'                                                 /
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

import (
	"log"
	"net"
	"strconv"
)

// Server base config interface to find server.
// This type is to be embedded or used in a config struct.
type Server struct {
	Host string
	Ports
}

// Default Server receives choice string for host.
func (srv *Server) Default(choice string) {
	// Default Host names
	switch choice {
	case "game":
		srv.Host = "davsk.sytes.net"
		break
	case "acct":
		srv.Host = "universe.gameserve.com"
		break
	default:
		srv.Host = choice
	}

	// Default Ports
	srv.Ports.Default()
}

// RpcConn returns a net.Conn to the RPC server.
// TODO(ds) This may not work on Unix as Unix wants a file system path.
func (srv *Server) RpcConn() net.Conn {
	conn, err := net.Dial("tcp", srv.Host+":"+strconv.Itoa(int(srv.Rpc)))
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	return conn
}
