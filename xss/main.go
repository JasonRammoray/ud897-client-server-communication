package main

import (
	"fmt"
	"bytes"
	"flag"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/GeertJohan/go.rice"
)

//go:generate rice embed-go
var (
	box = rice.MustFindBox("assets")
)

func main() {
	var (
		port = flag.Int("port", 8080, "Port to listen on")
	)
	flag.Parse()

	log.Printf("Running decoder server on decoder.127.0.0.1.xip.io:%d", *port)
	log.Printf("Running bad website on badwebsite.127.0.0.1.xip.io:%d", *port)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.Host, "decoder"):
			decodeServer(w, r)
		case strings.Contains(r.Host, "badwebsite"):
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
	if strings.HasSuffix("/", key) {
		key += "index.html"
	}
	fileContents, err := box.String(key)
	if err != nil {
		log.Printf("Could not find file %s", key)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tpl, err := template.New("").Parse(fileContents)
	if err != nil {
		log.Printf("Could not parse template %s: %s", key, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		log.Printf("Could not execute template %s: %s", key, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	io.Copy(w, buf)
}
