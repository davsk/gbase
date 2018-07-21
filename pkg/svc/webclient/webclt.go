// /////////////////////////////////////////////////////////////
// 'webclient.go'                                                 /
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

// Package webclient is a console client application
// for Universe 4.0 game.
//
// Overview
//
// webclient can be used alone in main to create
// a client application or it may be used with acctsrv
// and gameserver to create a standalone executable.
//
// Note that services must be started in sequence.
//    acct_srv.MustStart()
//    davsk_game_srv.MustStart()
//    webclient.MustStart()
//    nothing.Do()
package webclient

import (
	"davsk.net/gbase/pkg/must"
)

// Start starts the webclient service returns nil when ready
// or returns error on failure.
func Start() error {
	return nil
}

// MustStart starts the webclient service returns when ready
// and panics on failure.
func MustStart() {
	must.Do(Start())
}
