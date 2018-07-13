// /////////////////////////////////////////////////////////////
// 'toml_config_test.go'                                       /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 12, 2018                                               /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

package toml_config

import (
	"fmt"
	"testing"
)

const kTest = "unittest"

type ConfigInterface struct {
	Title string
}

func TestLoad(t *testing.T) {
	var v ConfigInterface
	err := Load(kTest, &v)
	fmt.Println(err)
	// Output:
	// open unittest.toml: The system cannot find the file specified.
}

func TestMustSave(t *testing.T) {
	var v ConfigInterface
	v.Title = kTest
	MustSave(kTest, &v)
}

func TestMustLoad(t *testing.T) {
	var v ConfigInterface
	MustLoad(kTest, &v)
	fmt.Println(v)
	// Output
	// {unittest}
}
