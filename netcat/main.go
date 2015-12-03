package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const successMessage = `HTTP/1.1 999 UDACITY EXERCISE SUCCES
Date: %s
Content-Type: application/udacity-exercise-token; charset=UTF-8
Content-Length: 49
Connection: close

Luckily, you don't need to do this stuff anymore!
`

var checks []func(r *http.Request) bool = []func(r *http.Request) bool{
	func(r *http.Request) bool {
		return r.Method == "OPTIONS"
	},
	func(r *http.Request) bool {
		return r.Header.Get("X-Udacity-Exercise-Header") != ""
	},
	func(r *http.Request) bool {
		return r.Header.Get("Date") == "Wed, 11 Jan 1995 23:00:00 GMT"
	},
}

func main() {
	log.Printf("Running webserver on netcat.127.0.0.1.xip.io:8080")

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, f := range checks {
			if !f(r) {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				log.Printf("Request metadata did not match expectations")
				return
			}
		}
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("Error preparing for hijack")
			return
		}
		c, buf, err := hj.Hijack()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("Error hijacking connection: %s", err)
			return
		}
		defer c.Close()
		defer buf.Flush()

		buf.Write([]byte(fmt.Sprintf(successMessage, time.Now().Format(time.RFC1123))))
	}))
	http.ListenAndServe(":8080", nil)
}
