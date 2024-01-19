package strategy

import "gnana997/load-balancer-go/pkg/domain"

type WeightedRoundRobinStrategy struct {
}

func (wrr *WeightedRoundRobinStrategy) Next(domains []*domain.Server) (*domain.Server, error) {
	return nil, nil
}
