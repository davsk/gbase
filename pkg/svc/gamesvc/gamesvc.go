// /////////////////////////////////////////////////////////////
// 'gamesvc.go'                                                /
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

// Package gamesvc starts an RPC microservice module
// for the Universe 4.0 game.
package gamesvc // import "davsk.net/gbase/pkg/gamesvc"

import (
	"database/sql"
	"log"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/must"
	_ "github.com/lib/pq"
)

// Start starts the gamesvc service returns nil when ready
// or returns error on failure.
func Start() error {
	// Load config.
	cfg := config.NewGameServer()

	// Open database.
	db, err := sql.Open("postgres",
		cfg.Game.ConnectionStr())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// handlers

	// start

	return nil
}

// MustStart starts the gamesvc service returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
