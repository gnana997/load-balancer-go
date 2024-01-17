package config

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig(strings.NewReader(`
services:
  - name: "Test"
    replicas:
      - "localhost:8081"
      - "localhost:8082"
strategy: "RoundRobin"
`))
	if err != nil {
		t.Errorf("Error should be nil but got %s", err)
	}
	if conf.Strategy != "RoundRobin" {
		t.Errorf("Expected strategy to be RoundRobin but got %s", conf.Strategy)
	}
	if len(conf.Services) != 1 {
		t.Errorf("Expected 1 service but got %d", len(conf.Services))
	}

	fmt.Printf("%+v\n", conf)
}
