// Copyright 2020 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/twitchtv/twirp-example/internal/hooks"
	"github.com/twitchtv/twirp-example/internal/server"
	"github.com/twitchtv/twirp-example/rpc/haberdasher"
)

func main() {
	hook := hooks.LoggingHooks(os.Stderr)
	twirpServer := server.NewHaberdasherServer()

	twirpHandler := haberdasher.NewHaberdasherServer(twirpServer, hook)
	mux := http.NewServeMux()
	mux.Handle(haberdasher.HaberdasherPathPrefix, twirpHandler)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, int64(10<<20)) // 10  MiB max per request
		mux.ServeHTTP(w, r)
	})

	// This port can be in a config file at some point
	port := 8080
	appServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	shutdownCh := make(chan bool)
	go func() {
		shutdownSIGch := make(chan os.Signal, 1)
		signal.Notify(shutdownSIGch, syscall.SIGINT, syscall.SIGTERM)

		<-shutdownSIGch

		log.Println("Received shutdown signal request")
		if err := appServer.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shutdown server: %s", err.Error())
		}
		close(shutdownCh)
	}()

	log.Printf("Starting example twirp service on port %d \n", port)
	if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to listen and serve: %s", err.Error())
	}

	<-shutdownCh
	log.Println("Server successfully shutdown")
}
