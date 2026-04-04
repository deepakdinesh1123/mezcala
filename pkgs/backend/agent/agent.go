package agent

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/moby/moby/client"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
)

type Agent struct {
	dc     *client.Client
	js     jetstream.JetStream
	logger *zerolog.Logger
}

func NewAgent(ctx context.Context, js jetstream.JetStream, logger *zerolog.Logger) (*Agent, error) {
	dc, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Agent{dc: dc, js: js, logger: logger}, nil
}

func (a *Agent) StartAsync(ctx context.Context) <-chan error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- a.Start(ctx)
		close(errCh)
	}()

	return errCh
}

func (a *Agent) Start(ctx context.Context) error {
	syncSub, err := a.js.CreateOrUpdateConsumer(
		ctx,
		"DB-SYNC",
		jetstream.ConsumerConfig{
			Name:          "sync_worker",
			Durable:       "sync_worker",
			MaxAckPending: 5,
		},
	)
	if err != nil {
		return err
	}

	adminSub, err := a.js.CreateOrUpdateConsumer(
		ctx,
		"DB-ADMIN",
		jetstream.ConsumerConfig{
			Name:          "admin_worker",
			Durable:       "admin_worker",
			MaxAckPending: 5,
		},
	)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Go(func() {
		a.runLoop(ctx, syncSub, a.handleSync)
	})

	wg.Go(func() {
		a.runLoop(ctx, adminSub, a.handleAdmin)
	})

	wg.Wait()

	return nil
}

func (a *Agent) runLoop(ctx context.Context, sub jetstream.Consumer, handler func(context.Context, jetstream.Msg) error) {
	a.logger.Debug().Msgf("starting consumer %s", sub.CachedInfo().Name)
	for {
		// Fetch blocks up to the timeout; use a short deadline so we can
		// check ctx.Done() frequently without spinning.
		msgs, err := sub.Fetch(1, jetstream.FetchMaxWait(2*time.Second))
		if err != nil {
			if errors.Is(err, nats.ErrTimeout) {
				// No messages yet — check if we should exit.
				if ctx.Err() != nil {
					return
				}
				continue
			}
			// Connection-level or subscription error — bail out.
			return
		}

		for msg := range msgs.Messages() {
			if err := handler(ctx, msg); err != nil {
				// msg.Nak()
				continue
			}
			msg.Ack()
		}

		// Check for cancellation after each batch.
		if ctx.Err() != nil {
			return
		}
	}
}
