// /////////////////////////////////////////////////////////////
// 'con_client.go'                                             /
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

// cmd con_client receives input from the user,
// communicates with the game_server,
// displays the results to the user.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/gbase/pkg/svc/conclient"
)

func main() {
	conclient.MustStart()
	nothing.MustDo()
}
