package main

import (
	"flag"
	"fmt"
	"gnana997/load-balancer-go/pkg/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	port = flag.Int("port", 8080, "Port to listen on")
)

type Balancer struct {
	Config      *config.Config
	ServerLists map[string]*config.ServerList
}

func NewBalancer(cfg *config.Config) *Balancer {
	serverLists := make(map[string]*config.ServerList)
	for _, service := range cfg.Services {
		serverList := &config.ServerList{
			Servers: make([]*config.Server, 0),
		}
		for _, replica := range service.Replicas {
			url, err := url.Parse(replica)
			if err != nil {
				log.Fatal(err)
			}
			serverList.Servers = append(serverList.Servers, &config.Server{
				Url:   url,
				Proxy: httputil.NewSingleHostReverseProxy(url),
			})
		}
		serverLists[service.Name] = serverList
	}

	return &Balancer{
		Config:      cfg,
		ServerLists: serverLists,
	}
}

func (b *Balancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO: Implement strategy per Service forwading
	fmt.Printf("Recieved request %s\n", r.Host)
	sl, ok := b.ServerLists["Test"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	next := sl.Next()
	fmt.Printf("Forwarding request to %s\n", sl.Servers[next].Url.Host)
	sl.Servers[next].Proxy.ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	services := make([]config.Service, 0)

	services = append(services, config.Service{
		Name: "Test",
		Replicas: []string{
			"http://127.0.0.1:8081",
			"http://127.0.0.1:8082",
		},
	})

	conf := &config.Config{
		Services: services,
		Strategy: "RoundRobin",
	}

	loadBalancer := NewBalancer(conf)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: loadBalancer,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
