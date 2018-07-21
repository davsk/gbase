// /////////////////////////////////////////////////////////////
// 'GameMaster.go'                                             /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 15, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// package GameMaster defines RPC interface of GameMaster
// for the gameserver for Universe 4.0.
package GameMaster

import "davsk.net/gbase/pkg/shared"

// GameMaster interface for the gameserver.
// List all RPC functions supported.
//   Usage: server.Register(GameMaster)
type GameMaster interface {
	Multiply(args *shared.Args, reply *int) error
	Divide(args *shared.Args, quo *shared.Quotient) error
}
