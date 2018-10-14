package sender

import dsender "github.com/go-zen-chu/peripush/pkg/domain/sender"

type factory struct {}

func NewFactory() dsender.Factory {
return &factory{}
}

func (f *factory) Service(pushGatewayURL string) dsender.Service {
	return NewSender(pushGatewayURL)
}