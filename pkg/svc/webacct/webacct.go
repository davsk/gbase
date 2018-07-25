// /////////////////////////////////////////////////////////////
// 'webacct.go'                                                /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 22, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Package webacct .
package webacct

import (
	"database/sql"

	"davsk.net/gbase/pkg/must"
)

type WebAcct struct {
	db *sql.DB
}

func (wa *WebAcct) Start(acctDb *sql.DB) (err error) {

	return err
}

func (wa *WebAcct) MustStart(acctDb *sql.DB) {
	must.Do(wa.Start(acctDb))
}
