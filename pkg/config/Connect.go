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

// Connect base config interface to connect.
// This type is to be embedded or used in a config struct.
type Connect struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

// Default Connect receives choice of databases.
func (c *Connect) Default(choice string) {
	switch choice {
	case "game":
		c.Host = "localhost"
		c.Port = 5432
		c.Database = "universe"
		c.User = "postgres"
		c.Password = "password"
		break
	case "acct":
		c.Host = "localhost"
		c.Port = 5432
		c.Database = "business"
		c.User = "postgres"
		c.Password = "password"
		break
	default:
		c.Host = "localhost"
		c.Port = 5432
		c.Database = choice
		c.User = "postgres"
		c.Password = "password"
	}
}
