// /////////////////////////////////////////////////////////////
// 'must.go'                                                   /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 20, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Package must provides configured error response for
// functions that must complete without error.
package must // import "davsk.net/gbase/pkg/must"

import "log"

// var opt is a local static variable containing
// module runtime configuration.
var opt option = option{0, nil}

// interface ErrorReporter implements the func Report(error).
// A user defined ErrorReporter may be used by this package.
type ErrorReporter interface {
	Report(error)
}

// type option is a private struct defining properties locally
// that may be configured at runtime.
type option struct {
	method int
	er     ErrorReporter
}

// SetMethod receives an int, sets opt.method, clears opt.er.
func SetMethod(x int) {
	opt.method = x
	opt.er = nil
}

// SetReporter receives an ErrorReporter, sets opt.er.
func SetReporter(er ErrorReporter) {
	opt.er = er
}

// func Do receives an error and handles it or not.
//    Usage:
//    must.Do(Something()) // reports error if Something() fails.
//    datum,err := LoadStuff()
//    must.Do(err) // reports error if LoadStuff() fails.
func Do(err error) {
	if err != nil {
		if opt.er != nil {
			opt.er.Report(err)
		} else {
			switch opt.method {
			case 0:
				log.Fatal(err)
			default:
				panic(err)
			}
		}
	}
}
