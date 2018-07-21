// /////////////////////////////////////////////////////////////
// 'guiclient.go'                                                 /
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

// Package guiclient is a console client application
// for Universe 4.0 game.
//
// Overview
//
// guiclient can be used alone in main to create
// a client application or it may be used with acctserver
// and gameserver to create a standalone executable.
//
// Note that services must be started in sequence.
//    acct_server.MustStart()
//    game_server.MustStart()
//    guiclient.MustStart()
//    nothing.Do()
package guiclient

import (
	"davsk.net/gbase/pkg/must"
)

// Start starts the guiclient service returns nil when ready
// or returns error on failure.
func Start() error {
	return nil
}

// MustStart starts the guiclient service returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
