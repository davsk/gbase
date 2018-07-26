// /////////////////////////////////////////////////////////////
// 'Ports.go'                                                  /
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

// Ports base config interface for server io.
// This type is to be embedded or used in a config struct.
type Ports struct {
	Http  uint16
	Https uint16
	Rpc   uint16
}

// Default Ports.
func (pt *Ports) Default() {
	pt.Http = 80
	pt.Https = 433
	pt.Rpc = 5000
}
