// /////////////////////////////////////////////////////////////
// 'conclt.go'                                                 /
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

// Package conclt is a console client service
// for Universe 4.0 game.
//
// Overview
//
// conclt can be used alone in main to create
// a client application or it may be used with acctsvc
// and gamesvc to create a standalone executable.
//
// Note that services must be started in sequence.
//    acctserver.MustStart()
//    gameserver.MustStart()
//    conclt.MustStart()
//    nothing.Do()
package conclt

import "log"

// Start starts the conclt service,
// returns nil when ready
// or returns error on failure.
func Start() error {
	return nil
}

// MustStart starts the ConClient returns when ready
// and panics on failure.
func MustStart() {
	if err := Start(); err != nil {
		log.Fatal(err)
	}
}
