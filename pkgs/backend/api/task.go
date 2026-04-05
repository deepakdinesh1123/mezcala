package api

import (
	"context"
	"fmt"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/nats-io/nats.go/jetstream"
)

func (s *Server) GetTaskUpdates(ctx context.Context, req spec.GetTaskUpdatesRequestObject) (spec.GetTaskUpdatesResponseObject, error) {
	_, err := s.js.CreateOrUpdateConsumer(
		ctx, fmt.Sprintf(spec.TASK_RESP_SUB, req.TaskID), jetstream.ConsumerConfig{
			Name:          "sync_worker",
			Durable:       "sync_worker",
			MaxAckPending: 5,
		},
	)
	if err != nil {

	}

	return spec.GetTaskUpdates200TexteventStreamResponse{}, nil
}
