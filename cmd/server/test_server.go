package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 8081, "Port to listen on")

type DemoServer struct{}

func (ds *DemoServer) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Hello"))
}

func main() {
	flag.Parse()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), &DemoServer{}); err != nil {
		log.Fatal(err)
	}
}
