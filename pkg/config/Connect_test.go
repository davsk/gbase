// /////////////////////////////////////////////////////////////
// 'Connect_test.go'                                           /
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

func TestConnect_Default(t *testing.T) {
	var game, acct, stuff Connect
	game.Default("game")
	acct.Default("acct")
	stuff.Default("stuff")
	fmt.Println(game, acct, stuff)
	// Output:
	// {localhost 5432 universe postgres password} {localhost 5432 business postgres password} {localhost 5432 stuff postgres password}
}
