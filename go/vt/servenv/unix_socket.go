// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package servenv

import (
	"flag"
	"net"
	"net/http"

	log "github.com/golang/glog"
)

var (
	socketFile = flag.String("socket_file", "", "Local unix socket file to listen on")
)

func serveSocketFile() {
	if *socketFile == "" {
		log.Infof("Not listening on socket file")
		return
	}

	l, err := net.Listen("unix", *socketFile)
	if err != nil {
		log.Fatalf("Error listening on socket file %v: %v", *socketFile, err)
	}
	log.Infof("Listening on socket file %v", *socketFile)
	go http.Serve(l, nil)
}
