package strategy

import (
	"fmt"
	"gnana997/load-balancer-go/pkg/domain"
	"sync"
)

type RoundRobinStrategy struct {
	mu      sync.Mutex
	current int
}

func (r *RoundRobinStrategy) Next(domains []*domain.Server) (*domain.Server, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	seen := 0
	var picked *domain.Server
	for seen < len(domains) {
		picked = domains[r.current]
		seen += 1
		r.current = (r.current + 1) % len(domains)
		if picked.IsAlive() {
			break
		}
	}
	if picked == nil || seen == len(domains) {
		return nil, fmt.Errorf("checked all the %d server. none of them are available", seen)
	}

	return picked, nil
}
