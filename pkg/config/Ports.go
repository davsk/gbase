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

import (
	"log"
	"net/http"
)

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

func (pt *Ports) Start(mux *http.ServeMux) (err error) {
	// redirect every http request to https
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	go http.ListenAndServeTLS(":443",
		"cert.pem",
		"key.pem",
		mux)

	// return status
	return err
}

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}

func index(w http.ResponseWriter, req *http.Request) {
	// all calls to unknown url paths should return 404
	if req.URL.Path != "/" {
		log.Printf("404: %s", req.URL.String())
		http.NotFound(w, req)
		return
	}
	http.ServeFile(w, req, "index.html")
}
