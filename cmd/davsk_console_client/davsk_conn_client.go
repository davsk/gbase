// /////////////////////////////////////////////////////////////
// 'davsk_console_client.go'                                   /
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

// cmd davsk_console_client receives input from the user,
// communicates with the davsk_game_srv,
// displays the results to the user.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/gbase/pkg/svc/consoleclient"
	"github.com/goinggo/tracelog"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)
	consoleclient.MustStart()
	nothing.MustDo()
	tracelog.Stop()
}
