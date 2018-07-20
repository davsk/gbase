// /////////////////////////////////////////////////////////////
// 'webclient.go'                                              /
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

// cmd webclient receives input from the user,
// communicates with the game server,
// displays the results to the user.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/gbase/pkg/svc/webclt"
)

func main() {
	webclt.MustStart()
	nothing.MustDo()
}
