////////////////////////////////////////////////////////////////
// '$name'                                                     /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
////////////////////////////////////////////////////////////////

// 'toml_config.go'

// by David Skinner
// on July 11, 2018
// for Davsk℠ Universe 4.0 project gbase

// package toml_config
package toml_config

import "davsk.net/gbase/pkg/custom_path"

const kToml = ".toml" // Config file type

func Load(title string) error {
	filePath := custom_path.Join(title + kToml)

	return nil
}

func Save(title string) error {
	filePath := custom_path.Join(title + kToml)

	return nil
}

func MustLoad(title string) {

}

func MustSave(title string) {

}
