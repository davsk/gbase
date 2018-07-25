// /////////////////////////////////////////////////////////////
// 'install.go'                                                /
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

// Package install performs download, config, setup,
// installation of Davsk℠ Universe 4.0 game and services.
package install

import (
	"davsk.net/gbase/pkg/must"
	"github.com/goinggo/tracelog"
)

const traceTitle = "install"

func Start() (err error) {
	// const FunctionName for tracelog.
	const traceFunctionName = "Start"
	tracelog.Started(traceTitle, traceFunctionName)

	// TODO(dls) Write code.

	// return status
	return err
}

func MustStart() {
	must.Do(Start())
}
