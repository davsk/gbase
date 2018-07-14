// /////////////////////////////////////////////////////////////
// 'Owner.go'                                                  /
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

// Owner base config interface for clients.
// This type is to be embedded or used in a config struct.
type Owner struct {
	Id    uint64 // 'toml:"OwnerId"'
	email string
}

// Default Owner.
func (o *Owner) Default() {
	o.Id = 0
	o.email = "email"
}
