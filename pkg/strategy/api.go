package strategy

import (
	"fmt"
	"gnana997/load-balancer-go/pkg/domain"
)

const (
	RoundRobin         = "RoundRobin"
	WeightedRoundRobin = "WeightedRoundRobin"
	Unknown            = "Unknown"
)

var strategies map[string]func() BalacingStrategy

type BalacingStrategy interface {
	Next(domains []*domain.Server) (*domain.Server, error)
}

type ServerList struct {
	Servers []*domain.Server

	Strategy BalacingStrategy
}

func init() {
	strategies = make(map[string]func() BalacingStrategy)
	strategies[RoundRobin] = func() BalacingStrategy {
		return &RoundRobinStrategy{
			Offset: uint32(0),
		}
	}
	strategies[WeightedRoundRobin] = func() BalacingStrategy {
		return &WeightedRoundRobinStrategy{}
	}
}

func LoadStrategy(Name string) BalacingStrategy {
	if strategy, ok := strategies[Name]; ok {
		fmt.Printf("Recieved Name is %s", Name)
		return strategy()
	}
	return strategies[RoundRobin]()
}
