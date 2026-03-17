package mq

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

func SetupJetStream(ctx context.Context, js jetstream.JetStream) error {
	_, err := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "DB-SYNC",
		Subjects: []string{"tasks.db_sync.*"},
		MaxAge:   2 * time.Hour,
	})
	if err != nil {
		fmt.Println("error is", err)
		return err
	}

	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "DB-ADMIN",
		Subjects: []string{"tasks.db_admin.*"},
		MaxAge:   1 * time.Hour,
	})
	if err != nil {
		fmt.Println("error is", err)
		return err
	}
	return nil
}
