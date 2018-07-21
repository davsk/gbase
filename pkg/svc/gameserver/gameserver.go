// /////////////////////////////////////////////////////////////
// 'gameserver.go'                                             /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 20, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Package gameserver starts an RPC microservice module
// for the Universe 4.0 game.
package gameserver // import "davsk.net/gbase/pkg/svc/gameserver"

import (
	"log"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/must"
)

// Start starts the gameserver service returns nil when ready
// or returns error on failure.
func Start() error {
	// Load config.
	log.Println("Loading config file.")
	cfg := config.NewGameSrv()

	// Open database.
	db := cfg.Game.MustOpenDatabase()
	defer db.Close()

	// handlers

	// start

	return nil
}

// MustStart starts the gameserver service returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
