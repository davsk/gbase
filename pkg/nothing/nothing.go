// /////////////////////////////////////////////////////////////
// 'nothing.go'                                                /
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

// Package nothing does nothing until you manually exit with
// CTRL-C.
//
// TODO(ds) Any suggestions on how to test this?
package nothing

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"davsk.net/gbase/pkg/must"
)

// Do nothing until you manually exit with CTRL-C.
func Do() error {
	fmt.Println("Waiting for you to manually exit with CTRL-C or")
	fmt.Println("Stop Process with Ctrl-F2 in JetBrains GoLand & WebStorm.")

	// wait
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // This will block until you manually exit with CTRL-C
	return nil
}

// MustDo nothing until you manually exit with CTRL-C,
// panics on error.
func MustDo() {
	must.Do(Do())
}
