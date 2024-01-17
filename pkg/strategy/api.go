package strategy

import (
	"gnana997/load-balancer-go/pkg/domain"
)

type FetchStrategy func(name string) BalacingStrategy

type BalacingStrategy interface {
	Next(domains []*domain.Server) (*domain.Server, error)
}

type ServerList struct {
	Servers []*domain.Server

	Strategy BalacingStrategy
}
