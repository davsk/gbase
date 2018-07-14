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
type Ports struct {
	Http  uint16
	Https uint16
	Rpc   uint16
}

// Default Ports.
func (p *Ports) Default() {
	p.Http = 80
	p.Https = 433
	p.Rpc = 5000
}
