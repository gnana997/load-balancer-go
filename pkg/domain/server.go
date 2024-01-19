package domain

import (
	"net/http/httputil"
	"net/url"
)

// Server is an instance of a running server
type Server struct {
	Url   *url.URL
	Proxy *httputil.ReverseProxy
}

type Replica struct {
	Host   string `yaml:"host"`
	Weight int    `yaml:"weight"`
}

type Service struct {
	Name string `yaml:"name"` // name of the serive
	// can be subdomain or regex
	Matcher  string    `yaml:"matcher"`  // prefix of the url to match the service
	Replicas []Replica `yaml:"replicas"` // replicas of the service like all the ips of the service
	Strategy string    `yaml:"strategy"` // name of the strategy for load balancing between the replicas
}
