// /////////////////////////////////////////////////////////////
// 'davsk_install.go'                                          /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 25, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

//  cmd davsk_install
package main

import (
	"davsk.net/gbase/internal/pkg/install"
	"github.com/goinggo/tracelog"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)
	install.MustStart()
	tracelog.Stop()
}
