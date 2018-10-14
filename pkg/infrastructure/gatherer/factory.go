package gatherer

import dgatherer "github.com/go-zen-chu/peripush/pkg/domain/gatherer"

type factory struct {

}

func NewFactory() dgatherer.Factory {
	return &factory{}
}

func (f *factory) Service() dgatherer.Service  {
	return NewService()
}