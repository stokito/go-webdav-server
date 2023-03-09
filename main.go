package main

import (
	"github.com/stokito/go-http-server-basic-auth"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
)

var credentials = map[string]string{
	"admin": "pass",
	"alice": "pass",
	"bob":   "pass",
}

func main() {
	dav := &webdav.Handler{
		FileSystem: webdav.Dir("./"),
		LockSystem: webdav.NewMemLS(),
		Logger:     davLog,
	}
	authHandler := basicauth.NewAuthHandlerWrapper(
		dav,
		credentials,
		"WebDAV",
		[]string{"/robots.txt", "/favicon.ico"},
	)
	recoverHandler := &basicauth.RecoveryHandlerWrapper{
		Handler:  authHandler,
		ErrorLog: log.Printf,
	}
	http.ListenAndServe(":8080", recoverHandler)
}

func davLog(r *http.Request, err error) {
	if err != nil {
		log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
		return
	}
	log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
}
