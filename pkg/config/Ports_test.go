// /////////////////////////////////////////////////////////////
// 'Ports_test.go'                                             /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 14, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

package config

import (
	"fmt"
	"testing"
)

func TestPorts_Default(t *testing.T) {
	var p Ports
	p.Default()
	fmt.Println(p)
	// Output:
	// {80 433 5000}
}
