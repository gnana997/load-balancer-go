package config

import (
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Service struct {
	Name     string   `yaml:"name"`     // name of the serive
	Replicas []string `yaml:"replicas"` // replicas of the service like all the ips of the service
}

// Config is a representation of the configuration given to load balancer from a config source
type Config struct {
	Services []Service `yaml:"services"`

	Strategy string `yaml:"strategy"` // Name of the strategy for load balancing between the replicas
}

// Server is an instance of a running server
type Server struct {
	Url   *url.URL
	Proxy *httputil.ReverseProxy
}

type ServerList struct {
	Servers []*Server

	// the offset is to forward the the request to.
	// the next server should be (offset + 1)% len(servers)
	offset uint32
}

func (sl *ServerList) Next() uint32 {
	nxt := atomic.AddUint32(&sl.offset, 1)
	return nxt % uint32(len(sl.Servers))
}
