package command

import (
	"context"
	"github.com/go-zen-chu/peripush/pkg/infrastructure/gatherer"
	"github.com/go-zen-chu/peripush/pkg/infrastructure/sender"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
)

type Command interface {
	Run() error
}

type config struct {
	pushGatewayURL string
}

type command struct{}

func NewCommand() Command {
	return &command{}
}

var (
	pgURL = kingpin.Arg("--push-gateway-url", "URL of push gateway").Required().Envar("PUSH_GATEWAY_URL").String()
	period = kingpin.Arg("--period", "Period to fetch metrics and push it to pushgateway.").Default("5m").Envar("PERIOD").Duration()
)

func (c *command) Run() error {
	cnf := &config{
		pushGatewayURL: *pgURL,
	}
	// context to handle cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := loop(ctx, *period, cnf)
	return err
}

func loop(ctx context.Context, period time.Duration, cnf *config) error {
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case t := <-ticker.C:
			err := periodicalFunc(t, cnf)
			if err != nil {
				return errors.Wrap(err, "Error while performing function periodically")
			}
		}
	}
	return nil
}

func periodicalFunc(t time.Time, cnf *config) error {
	// gather metrics
	gf := gatherer.NewFactory()
	g, err := gf.Service().Gather()
	if err != nil {
		return err
	}
	// send to push gateway
	sf := sender.NewFactory()
	err = sf.Service(cnf.pushGatewayURL).Send(g)
	if err != nil {
		return err
	}
	return nil
}