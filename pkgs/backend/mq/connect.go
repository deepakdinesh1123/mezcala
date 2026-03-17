package mq

import (
	"errors"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func StartEmbeddedNatsServer() (*nats.Conn, error) {
	opts := &server.Options{
		DontListen:      true,
		JetStream:       true,
		JetStreamDomain: "embedded",
		ServerName:      "mez_nats",
	}

	ns, err := server.NewServer(opts)
	if err != nil {
		return nil, err
	}
	go ns.Start()

	if !ns.ReadyForConnections(5 * time.Second) {
		return nil, errors.New("NATS server not ready")
	}

	nc, err := nats.Connect(nats.DefaultURL,
		nats.InProcessServer(ns),
	)
	if err != nil {
		return nil, err
	}
	return nc, nil
}
