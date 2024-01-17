package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

var (
	port = flag.Int("port", 8080, "Port to listen on")
)

type Service struct {
	name     string   // name of the serive
	replicas []string // replicas of the service like all the ips of the service
}

// Config is a representation of the configuration given to load balancer from a config source
type Config struct {
	Services []Service

	Strategy string // Name of the strategy for load balancing between the replicas
}

// Server is an instance of a running server
type Server struct {
	url   *url.URL
	proxy *httputil.ReverseProxy
}

type ServerList struct {
	Servers []*Server

	// the offset is to forward the the request to.
	// the next server should be (offset + 1)% len(servers)
	offset uint32
}

func (sl *ServerList) Next() uint32 {
	nxt := atomic.AddUint32(&sl.offset, 1)
	if nxt >= uint32(len(sl.Servers)) {
		nxt -= uint32(len(sl.Servers))
	}
	return nxt
}

type Balancer struct {
	Config      *Config
	ServerLists map[string]*ServerList
}

func NewBalancer(config *Config) *Balancer {
	serverLists := make(map[string]*ServerList)
	for _, service := range config.Services {
		serverList := &ServerList{
			Servers: make([]*Server, 0),
		}
		for _, replica := range service.replicas {
			url, err := url.Parse(replica)
			if err != nil {
				log.Fatal(err)
			}
			serverList.Servers = append(serverList.Servers, &Server{
				url:   url,
				proxy: httputil.NewSingleHostReverseProxy(url),
			})
		}
		serverLists[service.name] = serverList
	}

	return &Balancer{
		Config:      config,
		ServerLists: serverLists,
	}
}

func (b *Balancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	flag.Parse()

	conf := &Config{}

	loadBalancer := NewBalancer(conf)
	server := http.Server{
		Addr:    ":" + string(rune(*port)),
		Handler: loadBalancer,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
