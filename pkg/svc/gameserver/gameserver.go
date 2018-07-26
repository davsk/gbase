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
	"net/http"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/handler"
	"davsk.net/gbase/pkg/must"
	"github.com/goinggo/tracelog"
)

// const Title for tracelog
const traceTitle = "gameserver"

// Start starts the gameserver service returns nil when ready
// or returns error on failure.
func Start() error {
	// const FunctionName for tracelog
	const traceFunctionName = "Start"
	tracelog.Started(traceTitle, traceFunctionName)

	cfg := config.NewGameSrv()
	tracelog.Info(traceTitle, traceFunctionName,
		"GameSrv := %v", cfg)

	// Open database.
	db := cfg.Game.MustOpenDatabase()
	defer db.Close()
	tracelog.Info(traceTitle, traceFunctionName,
		"GameDatabase := %v", db)

	// handlers
	// serve index (and anything else) as https.
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)

	// start
	cfg.Start(mux)

	// Completed successfully.
	tracelog.Completed(traceTitle, traceFunctionName)

	return nil
}

// MustStart starts the gameserver service returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
