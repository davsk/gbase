// /////////////////////////////////////////////////////////////
// 'consoleclient.go'                                          /
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

// Package consoleclient is a console client service
// for Universe 4.0 game.
//
// Overview
//
// consoleclient can be used alone in main to create
// a client application or it may be used with acctsrv
// and gameserver to create a standalone executable.
//
// Note that services must be started in sequence.
//    acctsrv.MustStart()
//    gameserver.MustStart()
//    consoleclient.MustStart()
//    nothing.Do()
package consoleclient

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/handler"
	"davsk.net/gbase/pkg/must"
	"github.com/goinggo/tracelog"
	"gopkg.in/urfave/cli.v2"
)

// const Title for tracelog
const traceTitle = "consoleclient"

// Start starts the consoleclient service,
// returns nil when ready
// or returns error on failure.
func Start() error {
	// const FunctionName for tracelog.
	const traceFunctionName = "Start"
	tracelog.Started(traceTitle, traceFunctionName)

	// Load config.
	cfg := config.NewConsoleClient()
	tracelog.Info(traceTitle, traceFunctionName,
		"ConsoleClient := %v", cfg)

	// Completed successfully.
	tracelog.Completed(traceTitle, traceFunctionName)

	// handlers
	// serve index (and anything else) as https.
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Index)

	// start
	cfg.Start(mux)

	// go console.Start()
	app := new(cli.App)
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// return status
	return nil
}

// MustStart starts the consoleclient service, returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
