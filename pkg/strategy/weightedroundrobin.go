package strategy

import (
	"fmt"
	"gnana997/load-balancer-go/pkg/domain"
	"log"
	"sync"
)

type WeightedRoundRobinStrategy struct {
	mu    sync.Mutex
	count []int
	curr  int
}

func (wrr *WeightedRoundRobinStrategy) Next(domains []*domain.Server) (*domain.Server, error) {
	wrr.mu.Lock()
	defer wrr.mu.Unlock()

	seen := 0
	var picked *domain.Server
	for seen < len(domains) {
		if wrr.count == nil {
			wrr.count = make([]int, len(domains))
			wrr.curr = 0
		}

		picked = domains[wrr.curr]
		if !picked.IsAlive() {
			seen += 1
			wrr.count[wrr.curr] = 0
			wrr.curr = (wrr.curr + 1) % len(domains)
			continue
		}

		cap := picked.GetMetadataOrDefaultInt("weight", 1)

		if wrr.count[wrr.curr] <= cap {
			wrr.count[wrr.curr] += 1
			return picked, nil
		}

		wrr.count[wrr.curr] = 0
		wrr.curr = (wrr.curr + 1) % len(domains)
		log.Printf("curr: %d, count: %v", wrr.curr, wrr.count)
		log.Printf("Selected server is %s", domains[wrr.curr].Url.String())
	}
	if picked == nil || seen == len(domains) {
		return nil, fmt.Errorf("checked all the %d server. none of them are available", seen)
	}
	return picked, nil
}
