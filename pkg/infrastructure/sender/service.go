package sender

import (
	dsender "github.com/go-zen-chu/peripush/pkg/domain/sender"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/client_golang/prometheus"
)

type sender struct{
	pushGatewayURL string
}

func NewSender(pushGatewayURL string) dsender.Service {
	return &sender{
		pushGatewayURL: pushGatewayURL,
	}
}

func (s *sender) Send(g prometheus.Gatherer) error {
	pusher := push.New(s.pushGatewayURL, "peripush")
	err := pusher.Gatherer(g).Push()
	return err
}