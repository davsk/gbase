// /////////////////////////////////////////////////////////////
// 'conclient.go'                                              /
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

// Package conclient is a console client service
// for Universe 4.0 game.
//
// Overview
//
// conclient can be used alone in main to create
// a client application or it may be used with acctsrv
// and gameserver to create a standalone executable.
//
// Note that services must be started in sequence.
//    acctsrv.MustStart()
//    gameserver.MustStart()
//    conclient.MustStart()
//    nothing.Do()
package conclient

import (
	"fmt"
	"log"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/must"
)

// Start starts the conclient service,
// returns nil when ready
// or returns error on failure.
func Start() error {
	log.Println("Loading config file.")
	cfg := config.NewConsoleClient()
	fmt.Println(cfg)

	return nil
}

// MustStart starts the conclient service, returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
