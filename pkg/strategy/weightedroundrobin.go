package strategy

import (
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

	if wrr.count == nil {
		wrr.count = make([]int, len(domains))
		wrr.curr = 0
	}

	cap := domains[wrr.curr].GetMetadataOrDefaultInt("weight", 1)

	if wrr.count[wrr.curr] <= cap {
		wrr.count[wrr.curr] += 1
		return domains[wrr.curr], nil
	}

	wrr.count[wrr.curr] = 0
	wrr.curr = (wrr.curr + 1) % len(domains)
	log.Printf("curr: %d, count: %v", wrr.curr, wrr.count)
	log.Printf("Selected server is %s", domains[wrr.curr].Url.String())

	return domains[wrr.curr], nil
}
