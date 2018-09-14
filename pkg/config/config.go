// /////////////////////////////////////////////////////////////
// 'config.go'                                                 /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 12, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Package config loads config data from file
// or creates the file if it does not exist.
//
// Overview
//
// Types with a New method are meant to be used to load/save
// program config.
//
// Types without a New method are meant to be embedded or used
// in another config struct.
//
// If you make any changes in the config within the program,
// you must then use the MustUpdate method.
package config // import "davsk.net/gbase/pkg/config"

// Const traceTitle for tracelog.
const traceTitle = "config"
