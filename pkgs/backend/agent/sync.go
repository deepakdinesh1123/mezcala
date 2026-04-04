package agent

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"
)

func (a *Agent) SyncDatabase() {
}

func (a *Agent) handleSync(ctx context.Context, msg jetstream.Msg) error {
	return nil
}
