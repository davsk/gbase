// /////////////////////////////////////////////////////////////
// 'acctsrv.go'                                             /
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

// package acctsrv starts an RPC microservice module
// for the Universe 4.0 game.
package acctsrv // import "davsk.net/gbase/pkg/svc/acctsrv"

import (
	"log"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/must"
)

func Start() error {
	// Load config.
	log.Println("Loading config file.")
	cfg := config.NewAcctSrv()

	// Open database.
	log.Println("Opening Acct Database.")
	db := cfg.Acct.MustOpenDatabase()
	defer db.Close()

	// handlers

	// start

	return nil
}

func MustStart() {
	must.Do(Start())
}
