package strategy

import (
	"gnana997/load-balancer-go/pkg/domain"
	"sync/atomic"
)

type RoundRobin struct {
	Offset uint32
}

func (r *RoundRobin) Next(domains []*domain.Server) (*domain.Server, error) {
	nxt := atomic.AddUint32(&r.Offset, 1)
	return domains[nxt%uint32(len(domains))], nil
}
