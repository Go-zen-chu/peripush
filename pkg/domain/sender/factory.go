package sender

package gatherer

type Factory interface {
	Service(pushGatewayURL string) Service
}
