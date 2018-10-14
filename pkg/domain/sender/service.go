package sender

import "github.com/prometheus/client_golang/prometheus"

type Service interface {
	Send(g prometheus.Gatherer) error
}

