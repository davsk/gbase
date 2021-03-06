// /////////////////////////////////////////////////////////////
// 'TestSrv_test.go'                                        /
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
	"fmt"
	"testing"
)

func TestNewTestServer(t *testing.T) {
	ts := NewTestSrv()
	fmt.Println(ts)
	// Output:
	// {config_test_server {80 433 5000} {8000 8080 5001} {localhost 5432 universe postgres password} {localhost 5432 business postgres password}}
}
