package health

import (
	"fmt"
	"gnana997/load-balancer-go/pkg/config"
	"gnana997/load-balancer-go/pkg/domain"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

type HealthChecker struct {
	servers []*domain.Server
	period  int
}

func NewChecker(conf *config.Config, servers []*domain.Server) (*HealthChecker, error) {
	if len(servers) == 0 {
		return nil, fmt.Errorf("no servers provided")
	}

	return &HealthChecker{
		servers: servers,
		period:  1,
	}, nil
}

func (hc *HealthChecker) Start() {
	log.Info("starting health checker")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		for _, server := range hc.servers {
			go hc.checkhealth(server)
		}
	}
}

func (hc *HealthChecker) checkhealth(server *domain.Server) {
	_, err := net.DialTimeout("tcp", server.Url.Host, 5*time.Second)
	if err != nil {
		log.Errorf("could not connect to server '%s'", server.Url.Host)
		old := server.SetLiveness(false)
		if old {
			log.Warnf("server '%s' is down", server.Url.Host)
		}
		return
	}
	old := server.SetLiveness(true)
	if !old {
		log.Infof("server '%s' is up", server.Url.Host)
	}
}
