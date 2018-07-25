// /////////////////////////////////////////////////////////////
// 'GameMaster/client.go'                                      /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 15, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Package GameMaster/Client implements RPC Client GameMaster
// interface for use by Clients in Universe 4.0.
package Client

import (
	"log"
	"net/rpc"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/shared"
)

// Create a struct, that mimics all methods provided by interface.
// It is not compulsory, we are doing it here,
// just to simulate a traditional method call.
type Game struct {
	client *rpc.Client
}

// New receives config.Server and returns Acct struct.
func New(server config.Server) *Game {
	return &Game{client: rpc.NewClient(server.RpcConn())}
}

// Divide is a test function.
func (t *Game) Divide(a, b int) shared.Quotient {
	args := &shared.Args{a, b}
	var reply shared.Quotient
	err := t.client.Call("Game.Divide", args, &reply)
	if err != nil {
		log.Fatal("acct error:", err)
	}
	return reply
}

// Multiply is a test function.
func (t *Game) Multiply(a, b int) int {
	args := &shared.Args{a, b}
	var reply int
	err := t.client.Call("Game.Multiply", args, &reply)
	if err != nil {
		log.Fatal("acct error:", err)
	}
	return reply
}
