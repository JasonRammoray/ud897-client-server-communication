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
	http.Handle("/", http.FileServer(box.HTTPBox()))
	http.ListenAndServe(":8080", nil)
}
