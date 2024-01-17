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
    matcher: "/api/v1"
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
	if conf.Services[0].Matcher != "/api/v1" {
		t.Errorf("Expected matcher to be /api/v1 but got %s", conf.Services[0].Matcher)
	}
	if conf.Services[0].Name != "Test" {
		t.Errorf("Expected name to be Test but got %s", conf.Services[0].Name)
	}
	if len(conf.Services[0].Replicas) != 2 {
		t.Errorf("Expected 2 replicas but got %d", len(conf.Services[0].Replicas))
	}

	fmt.Printf("%+v\n", conf)
}
