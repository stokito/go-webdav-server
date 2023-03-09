package main

import (
	"golang.org/x/net/webdav"
	"log"
	"net/http"
)

func main() {
	dav := &webdav.Handler{
		FileSystem: webdav.Dir("./"),
		LockSystem: webdav.NewMemLS(),
		Logger:     davLog,
	}
	http.ListenAndServe(":8080", dav)
}

func davLog(r *http.Request, err error) {
	if err != nil {
		log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
		return
	}
	log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
}
