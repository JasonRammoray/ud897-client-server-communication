// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/udacity/ud897-client-server-communication/utils"
)

//go:generate rice embed-go
var box *rice.Box

func main() {
	box = rice.MustFindBox("assets")
	var (
		port = flag.Int("port", 8080, "Port to listen on")
	)
	flag.Parse()

	log.Printf("Running decoder server on decoder.127.0.0.1.xip.io:%d", *port)
	log.Printf("Running bad website on badwebsite.127.0.0.1.xip.io:%d", *port)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname := strings.Split(r.Host, ":")[0]
		switch {
		case strings.HasPrefix(hostname, "decoder."):
			decodeServer(w, r)
		case strings.HasPrefix(hostname, "badwebsite."):
			badWebsite(w, r)
		}
	}))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Could not start webserver on :%d: %s", *port, err)
	}
}

var (
	secret = "dAZxAgQAAHY="
)

func decodeServer(w http.ResponseWriter, r *http.Request) {
	x := []byte(r.FormValue("key"))
	y, _ := base64.StdEncoding.DecodeString(secret)
	if len(x) > len(y) {
		x = x[0:len(y)]
	}
	for i := range x {
		x[i] = x[i] ^ y[i]
	}
	log.Printf("Result: %s", x)
}

func badWebsite(w http.ResponseWriter, r *http.Request) {
	// Set the cookie that is supposed to be stolen
	http.SetCookie(w, &http.Cookie{
		Name:  "SESSION_ID",
		Value: "DEADBEEF",
	})

	// Disable XSS protection because securitylol
	w.Header().Set("X-XSS-Protection", "0")

	data := r.URL.Query()
	if _, ok := data["name"]; !ok {
		data["name"] = []string{"Anonymous"}
	}

	key := r.URL.Path[1:]
	err := utils.ExecuteTemplateInBox(w, box, key, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error executing template: %s", err)
		return
	}
}
