package main

import (
	"go-webdav-serv/webdav2"
	_ "golang.org/x/net/webdav"
	"log"
	"net/http"
)

func main() {
	dav := &webdav2.Handler{
		FileSystem: webdav2.Dir("./"),
		LockSystem: webdav2.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
			} else {
				log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
			}
		},
	}
	port := ":8090"
	println("Started on ", port)
	http.ListenAndServe(port, dav)
}
