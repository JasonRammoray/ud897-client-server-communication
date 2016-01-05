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

const successToken = "ba16d08e117ea066a43cf1332d78b0324b993cdcc0bd72ea8827f26d3457a32e"

var checks []func(r *http.Request) string = []func(r *http.Request) string{
	func(r *http.Request) string {
		if r.Method == "UDACITY" {
			return ""
		}
		return "8617c0f6c64d3d1a16858bc61236cd02e9ab7a3bb21df3dc796944eaa6f3bc"
	},
	func(r *http.Request) string {
		if r.Header.Get("X-Udacity-Exercise-Header") != "" {
			return ""
		}
		return "3627cdf4dae75dff97584d8d84d3c8821335d23ee1d57cfbe0be2b1bd22a45"
	},
	func(r *http.Request) string {
		if r.Header.Get("Date") == "Wed, 11 Jan 1995 23:00:00 GMT" {
			return ""
		}
		return "8053172ec777c51169271a3fa49644be6a59eb174828e98c2fcb628beabb"
	},
}

func main() {
	log.Printf("Running webserver on netcat.127.0.0.1.xip.io:8080")

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := successToken
		for _, f := range checks {
			if x := f(r); x!= "" {
				w.Header().Set("X-No-Success", "true");
				token = x
				break;
			}
		}

		w.Write([]byte(token))
		w.Write([]byte("\n"))
	}))
	http.ListenAndServe(":8080", nil)
}
