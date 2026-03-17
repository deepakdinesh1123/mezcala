package agent

import (
	"github.com/nats-io/nats.go/jetstream"
)

func (a *Agent) SyncDatabase() {
}

func (a *Agent) handleSync(jetstream.Msg) error {
	return nil
}
