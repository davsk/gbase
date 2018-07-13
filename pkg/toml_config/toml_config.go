// /////////////////////////////////////////////////////////////
// 'toml_config.go'                                            /
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

// package toml_config provides helper functions
// to load and save toml config files.
package toml_config

import (
	"log"
	"os"

	"davsk.net/gbase/pkg/custom_path"
	"github.com/BurntSushi/toml"
)

// Config file type extension.
const kToml = ".toml"

// Load receives required title string and struct interface
// and builds file name path from title
// and returns error code after attempting to fill interface.
func Load(title string, v interface{}) error {
	// create file path to config file.
	fpath := custom_path.Join(title + kToml)

	// attempt to read file.
	_, err := toml.DecodeFile(fpath, v)

	// return error code, nil if successful.
	return err
}

// Save receives required title string and struct interface
// and builds file name path from title
// and returns error code after attempting to save interface.
func Save(title string, v interface{}) error {
	// create file path to config file.
	fpath := custom_path.Join(title + kToml)

	// create config file.
	f, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enc := toml.NewEncoder(f) // create encoder.
	err = enc.Encode(v)       // write config file.

	// return error code, nil if successful.
	return err
}

// Load receives required title string and struct interface
// and builds file name path from title
// and after attempting to fill interface panics on failure.
func MustLoad(title string, v interface{}) {
	if err := Load(title, v); err != nil {
		panic(err)
	}
}

// Save receives required title string and struct interface
// and builds file name path from title
// and after attempting to save interface will panic on failure.
func MustSave(title string, v interface{}) {
	if err := Save(title, v); err != nil {
		panic(err)
	}
}
