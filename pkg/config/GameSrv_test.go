// /////////////////////////////////////////////////////////////
// 'GameSrv_test.go'                                        /
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

func TestNewGameServer(t *testing.T) {
	gs := NewGameSrv()
	fmt.Println(gs)
	// Output:
	// {config_game_server {80 433 5000} {universe.gameserve.com {80 433 5000}} {localhost 5432 universe postgres password}}
}
