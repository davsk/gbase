// /////////////////////////////////////////////////////////////
// 'Connect.go'                                                /
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
	"database/sql"
	"strconv"

	"davsk.net/gbase/pkg/must"
	_ "github.com/lib/pq"
)

// Connect base config interface to connect.
// This type is to be embedded or used in a config struct.
type Connect struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

// ConnectionStr returns a string from the Connect type.
func (cn *Connect) ConnectionStr() string {
	var cns string

	if len(cn.Host) > 0 {
		cns += "host=" + cn.Host + " "
	}
	if cn.Port > 0 {
		cns += "port=" + strconv.Itoa(int(cn.Port)) + " "
	}
	if len(cn.Database) > 0 {
		cns += "database=" + cn.Database + " "
	}
	if len(cn.User) > 0 {
		cns += "user=" + cn.User + " "
	}
	if len(cn.Password) > 0 {
		cns += "password=" + cn.Password
	}

	return cns
}

// OpenDatabase returns pointer to pointer to an sql.DB
// and error.
func (cn *Connect) OpenDatabase() (*sql.DB, error) {
	return sql.Open("postgres",
		cn.ConnectionStr())
}

// MustOpenDatabase returns pointer to an sql.DB,
// panics on fail.
//
// Usage:
//    db := cn.MustOpenDatabase()
//    defer db.close()
func (cn *Connect) MustOpenDatabase() *sql.DB {
	db, err := cn.OpenDatabase()
	must.Do(err)

	return db
}

// Default Connect receives choice of databases.
func (cn *Connect) Default(choice string) {
	switch choice {
	case "game":
		cn.Host = "localhost"
		cn.Port = 5432
		cn.Database = "universe"
		cn.User = "postgres"
		cn.Password = "password"
		break
	case "acct":
		cn.Host = "localhost"
		cn.Port = 5432
		cn.Database = "business"
		cn.User = "postgres"
		cn.Password = "password"
		break
	default:
		cn.Host = "localhost"
		cn.Port = 5432
		cn.Database = choice
		cn.User = "postgres"
		cn.Password = "password"
	}
}
