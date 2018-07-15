// /////////////////////////////////////////////////////////////
// 'AcctMaster.go'                                             /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on 7 15, 2018                                               /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// package AcctMaster defines RPC interface of AcctMaster
// for the AcctServer for Universe 4.0.
package AcctMaster

import "davsk.net/gbase/pkg/shared"

// AcctMaster interface for the AcctServer.
// List all RPC functions supported.
//   Usage: server.Register(AcctMaster)
type AcctMaster interface {
	Multiply(args *shared.Args, reply *int) error
	Divide(args *shared.Args, quo *shared.Quotient) error
}
