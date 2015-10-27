package utils

import (
	"log"
	"net/http"
)

type Logserver struct{}

func (l Logserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)
	for k, v := range r.Header {
		log.Printf("%s: %#v", k, v)
	}
	log.Printf("=== END OF REQUEST DATA ===")
}
