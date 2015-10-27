package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/GeertJohan/go.rice"
	"github.com/udacity/ud897-client-server-communication/utils"
)

//go:generate rice embed-go
var (
	box       = rice.MustFindBox("assets")
	logserver = utils.Logserver{}
)

func main() {
	log.Printf("Running logging server on logger.127.0.0.1.xip.io:8080")
	log.Printf("Running bad website on badwebsite.127.0.0.1.xip.io:8080")

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.Split(r.Host, ":")[0] {
		case "logger.127.0.0.1.xip.io":
			logserver.ServeHTTP(w, r)
		case "badwebsite.127.0.0.1.xip.io":
			badWebsite(w, r)
		}
	}))
	http.ListenAndServe(":8080", nil)
}

func badWebsite(w http.ResponseWriter, r *http.Request) {
	// Set the cookie that is supposed to be stolen
	http.SetCookie(w, &http.Cookie{
		Name:  "SESSION_ID",
		Value: "3735928559",
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
