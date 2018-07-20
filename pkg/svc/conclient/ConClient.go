// /////////////////////////////////////////////////////////////
// 'ConClient.go'                                              /
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

// package ConClient is a console client application
// for Universe 4.0 game.
//
// Overview
//
// ConClient can be used alone in main to create
// a client application or it may be used with AcctServer
// and GameServer to create a standalone executable.
//
// Note that services must be started in sequence.
//    acctserver.MustStart()
//    gameserver.MustStart()
//    conclient.MustStart()
//    nothing.Do()
package conclient

import "log"

// Start starts the ConClient returns nil when ready
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
