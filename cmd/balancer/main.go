package main

import (
	"flag"
	"fmt"
	"gnana997/load-balancer-go/pkg/config"
	"gnana997/load-balancer-go/pkg/domain"
	"gnana997/load-balancer-go/pkg/strategy"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var (
	port       = flag.Int("port", 8080, "Port to listen on")
	confgiFile = flag.String("config-path", "", "Path to config file")
)

type Balancer struct {
	Config      *config.Config
	ServerLists map[string]*strategy.ServerList
}

func NewBalancer(cfg *config.Config) *Balancer {
	// TODO: prevent multiple or invalid matchers before creating the server list
	serverLists := make(map[string]*strategy.ServerList)
	for _, service := range cfg.Services {
		balancingStrategy := strategy.FetchStrategy(func(name string) strategy.BalacingStrategy {
			fmt.Printf("Fetching strategy %s\n", name)
			return &strategy.RoundRobin{
				Offset: 0,
			}
		})(service.Strategy)

		serverList := &strategy.ServerList{
			Servers:  make([]*domain.Server, 0),
			Strategy: balancingStrategy,
		}
		for _, replica := range service.Replicas {
			url, err := url.Parse(replica)
			if err != nil {
				log.Fatal(err)
			}
			serverList.Servers = append(serverList.Servers, &domain.Server{
				Url:   url,
				Proxy: httputil.NewSingleHostReverseProxy(url),
			})
		}
		serverLists[service.Matcher] = serverList
	}

	return &Balancer{
		Config:      cfg,
		ServerLists: serverLists,
	}
}

func (b *Balancer) findServiceList(path string) (*strategy.ServerList, error) {
	fmt.Printf("Finding service for path %s\n", path)
	for matcher, serverList := range b.ServerLists {
		if strings.HasPrefix(path, matcher) {
			return serverList, nil
		}
	}
	return nil, fmt.Errorf("no service found for path %s", path)
}

func (b *Balancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO: Implement strategy per Service forwading
	fmt.Printf("Recieved request %s\n", r.Host)
	sl, err := b.findServiceList(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Not Found Error: %s", err)))
		return
	}
	next, err := sl.Strategy.Next(sl.Servers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error occured while choosing server: %s", err)))
		return
	}
	fmt.Printf("Forwarding request to %s\n", next.Url.Host)
	next.Proxy.ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	file, err := os.Open(*confgiFile)
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.LoadConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", config)

	conf := config

	loadBalancer := NewBalancer(conf)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: loadBalancer,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
