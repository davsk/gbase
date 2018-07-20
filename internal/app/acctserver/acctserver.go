// /////////////////////////////////////////////////////////////
// 'acctsvc.go'                                             /
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

// cmd acctsvc is an http/https/rpc server
// that accesses a database of accounts, users, and servers.
package main

import (
	"davsk.net/gbase/pkg/acctsvc"
	"davsk.net/gbase/pkg/nothing"
	_ "github.com/lib/pq"
)

const kVersion = "v0.1.0"

func main() {
	acctsvc.MustStart()
	nothing.MustDo()
}
