package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/udacity/ud897-client-server-communication/utils"
)

func main() {
	os.Mkdir("website", os.FileMode(0755))
	log.Printf("Created a folder called \"website\", all its contents")
	log.Printf("will be served on website.127.0.0.1.xip.io:8080")
	log.Printf("Running CORS website on cors.127.0.0.1.xip.io:8080")
	log.Printf("Running non-CORS website on noncors.127.0.0.1.xip.io:8080")

	website := http.FileServer(http.Dir("website"))
	logger := utils.Logserver{}
	corsLogger := utils.EnableCORS(logger)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname := strings.Split(r.Host, ":")[0]
		switch {
		case strings.HasPrefix(hostname, "website."):
			website.ServeHTTP(w, r)
		case strings.HasPrefix(hostname, "noncors."):
			logger.ServeHTTP(w, r)
		case strings.HasPrefix(hostname, "cors."):
			corsLogger.ServeHTTP(w, r)
		}
	}))
	http.ListenAndServe(":8080", nil)
}
