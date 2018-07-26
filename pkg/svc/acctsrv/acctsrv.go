// /////////////////////////////////////////////////////////////
// 'acctsrv.go'                                             /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 19, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Package acctsrv starts an RPC microservice module
// for the Universe 4.0 game.
package acctsrv // import "davsk.net/gbase/pkg/svc/acctsrv"

import (
	"net/http"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/handler"
	"davsk.net/gbase/pkg/must"
	"github.com/goinggo/tracelog"
)

const traceTitle = "acctsrv"

func Start() error {
	// const FunctionName for tracelog.
	const traceFunctionName = "Start"
	tracelog.Started(traceTitle, traceFunctionName)

	// Load config.
	cfg := config.NewAcctSrv()
	tracelog.Info(traceTitle, traceFunctionName,
		"AcctSrv := %v", cfg)

	// Open database.
	db := cfg.Acct.MustOpenDatabase()
	defer db.Close()
	tracelog.Info(traceTitle, traceFunctionName,
		"AcctDatabase := %v", db)

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

func MustStart() {
	must.Do(Start())
}
