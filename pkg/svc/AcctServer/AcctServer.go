// /////////////////////////////////////////////////////////////
// 'AcctServer.go'                                             /
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

// package AcctServer starts an RPC microservice module
// for the Universe 4.0 game.
package AcctServer // import "davsk.net/gbase/pkg/AcctServer"

import (
	"database/sql"
	"log"

	"davsk.net/gbase/pkg/config"
	"davsk.net/gbase/pkg/must"
	_ "github.com/lib/pq"
)

func Start() error {
	// Load config.
	cfg := config.NewAcctServer()

	// Open database.
	db, err := sql.Open("postgres",
		cfg.Acct.ConnectionStr())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// handlers

	// start

	return nil
}

func MustStart() {
	must.Do(Start())
}
