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
	"log"
	"net/http"
)

// Tokens have been generated with this JS snippet:
//
// crypto.subtle.digest("SHA-256", new TextEncoder("utf-8").encode(new Date().getTime().toString()))
//   .then(x => new Uint8Array(x))
//   .then(x => Array.from(x).map(x => x.toString(16)))
//   .then(x => x.join(''))
//   .then(x => console.log(x))

const successToken = "ba16d"

var checks []func(r *http.Request) string = []func(r *http.Request) string{
	func(r *http.Request) string {
		if r.Method == "UDACITY" {
			return ""
		}
		return "Error: Not the right HTTP method"
	},
	func(r *http.Request) string {
		if r.Header.Get("X-Udacity-Exercise-Header") != "" {
			return ""
		}
		return "Error: X-Udacity-Exercise-Header missing"
	},
	func(r *http.Request) string {
		if r.Header.Get("Date") == "Wed, 11 Jan 1995 23:00:00 GMT" {
			return ""
		}
		return "Error: Wrong Date"
	},
}

func main() {
	log.Printf("Running webserver on http://netcat.127.0.0.1.xip.io:8080")

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := successToken
		for _, f := range checks {
			if x := f(r); x != "" {
				w.Header().Set("X-No-Success", "true")
				token = x
				break
			}
		}

		w.Write([]byte(token))
		w.Write([]byte("\n"))
	}))
	http.ListenAndServe(":8080", nil)
}
