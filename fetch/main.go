package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

//go:generate rice embed-go
var (
	box = rice.MustFindBox("assets")
)

func main() {
	log.Printf("Running website on 127.0.0.1.xip.io:8080")
	http.Handle("/password.txt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && r.Header.Get("X-Udacity-Exercise") != "" {
			w.Write([]byte("Password: piquizahhai5aeh2fah9Uk"))
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	}))
	http.Handle("/", http.FileServer(box.HTTPBox()))
	http.ListenAndServe(":8080", nil)
}
