// /////////////////////////////////////////////////////////////
// 'test.go'                                                   /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 21, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Command test runs game on local test server
// using logfiles from consoleclient.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/gbase/pkg/svc/acctsrv"
	"davsk.net/gbase/pkg/svc/consoleclient"
	"davsk.net/gbase/pkg/svc/gameserver"
	"github.com/goinggo/tracelog"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)
	acctsrv.MustStart()
	gameserver.MustStart()
	consoleclient.MustStart()
	nothing.Do()
	tracelog.Stop()
}
