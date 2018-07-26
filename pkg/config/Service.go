// /////////////////////////////////////////////////////////////
// 'Service.go'                                                /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 25, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

package config

import (
	"log"
	"net/http"
	"strconv"

	"github.com/goinggo/tracelog"
)

// Type Service embeds Cert and Port.
type Service struct {
	Cert
	Ports
}

// Function Default receives name string and sets default values.
func (sv *Service) Default(name string) {
	sv.Cert.Default(name)
	sv.Ports.Default()
}

// Function Start
func (sv *Service) Start(mux *http.ServeMux) (err error) {
	// const FunctionName for tracelog.
	const traceFunctionName = "Start"
	tracelog.Started(traceTitle, traceFunctionName)

	// Only redirect if port is assigned.
	if sv.Http > 0 {
		// test for valid redirect
		if sv.Http != 80 || sv.Https != 443 {
			log.Fatal("Current redirect only supports" +
				" HTTP port 80 and HTTPS port 443.")
		}

		// redirect every http request to https
		go http.ListenAndServe(
			":"+strconv.Itoa(int(sv.Http)),
			http.HandlerFunc(redirect))
	}

	sv.Validate()

	go http.ListenAndServeTLS(
		":"+strconv.Itoa(int(sv.Https)),
		sv.CertFile,
		sv.KeyFile,
		mux)

	// return status
	return err
}

// Function redirect is a handler that redirects a HTTP Request
// to HTTPS.
//
// TODO(dls) Make it work when HTTP is not port 80 and/or HTTPS is not port 443.
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
