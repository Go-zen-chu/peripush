package gatherer

import "github.com/prometheus/client_golang/prometheus"

type Service interface {
	Gather() (prometheus.Gatherer, error)
}
