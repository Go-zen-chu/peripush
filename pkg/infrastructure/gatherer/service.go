package gatherer

import (
	"github.com/prometheus/client_golang/prometheus"
	dgatherer "github.com/go-zen-chu/peripush/pkg/domain/gatherer"
)

type service struct {

}

func NewService() dgatherer.Service {
	return &service{}
}

func (s *service) Gather() (prometheus.Gatherer, error) {
	rg := prometheus.NewRegistry()
	// register collector here
	rg.MustRegister()
	return rg, nil
}