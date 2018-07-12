////////////////////////////////////////////////////////////////
// 'custom_path_test.go'                                       /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 12, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
////////////////////////////////////////////////////////////////

package custom_path

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	fpath := Join("test.tst")
	fmt.Printf("%s", fpath)
}
