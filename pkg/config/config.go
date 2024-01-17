package config

import (
	"gnana997/load-balancer-go/pkg/domain"
)

// Config is a representation of the configuration given to load balancer from a config source
type Config struct {
	Services []domain.Service `yaml:"services"`

	Strategy string `yaml:"defaultStrategy"` // Name of the strategy for load balancing between the replicas
}
