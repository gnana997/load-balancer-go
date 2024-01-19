package domain

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strconv"
)

// Server is an instance of a running server
type Server struct {
	Url      *url.URL
	Proxy    *httputil.ReverseProxy
	Metadata map[string]string
}

func (s *Server) GetMetadataOrDefault(key string, defaultValue string) string {
	if s.Metadata != nil {
		if value, ok := s.Metadata[key]; ok {
			return value
		}
	}
	return defaultValue
}

func (s *Server) GetMetadataOrDefaultInt(key string, defaultValue int) int {
	if s.Metadata != nil {
		value := s.GetMetadataOrDefault(key, fmt.Sprintf("%d", defaultValue))
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return defaultValue
}
