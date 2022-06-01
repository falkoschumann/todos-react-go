package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"todos_backend_server/spa"
)

func main() {
	host, port := parseFlags()
	createRouter()
	runServer(host, port)
}

func parseFlags() (host string, port uint) {
	flag.StringVar(&host, "host", "", "the server listen on this `host` (default all)")
	flag.UintVar(&port, "port", 8080, "the server listen on this `port`")
	flag.Parse()
	return
}

func createRouter() {
	http.Handle("/", spa.Handler{})
}

func runServer(host string, port uint) {
	log.Println("Server listening on port", port)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
