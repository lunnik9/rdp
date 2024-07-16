package main

import (
	"github.com/kulaginds/rdp-html5/handler"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	if err := startServer(); err != nil {
		log.Fatalln(err)
	}
}

func startServer() error {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/connect", handler.Connect)

	log.Println("start web-server on :8080")

	return http.ListenAndServe(":8080", nil)
}
