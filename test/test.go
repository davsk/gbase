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

// cmd test runs game on local test server
// using logfiles from conclient.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/gbase/pkg/svc/acctserver"
	"davsk.net/gbase/pkg/svc/conclient"
	"davsk.net/gbase/pkg/svc/gameserver"
)

func main() {
	acctserver.MustStart()
	gameserver.MustStart()
	conclient.MustStart()
	nothing.Do()
}
