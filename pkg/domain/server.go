package domain

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"
)

// Server is an instance of a running server
type Server struct {
	Url      *url.URL
	Proxy    *httputil.ReverseProxy
	Metadata map[string]string

	mu    sync.RWMutex
	alive bool
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

func (s *Server) SetLiveness(alive bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	old := s.alive
	s.alive = alive
	return old
}

func (s *Server) IsAlive() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.alive
}
