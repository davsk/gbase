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

// cmd gameserver is an http/rpc server
// that accesses a game database and rpcs the acct_server.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/gbase/pkg/svc/gamesvc"
)

func main() {
	gamesvc.MustStart()
	nothing.MustDo()
}
