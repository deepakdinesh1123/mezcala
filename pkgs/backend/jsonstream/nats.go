package jsonstream

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
)

// NatsWriter implements io.Writer, publishing each write as an AgentResponse to NATS.
type NatsWriter struct {
	ctx    context.Context
	js     jetstream.JetStream
	taskID string
	logger *zerolog.Logger
}

func NewNatsWriter(ctx context.Context, js jetstream.JetStream, taskID string, logger *zerolog.Logger) *NatsWriter {
	return &NatsWriter{
		ctx:    ctx,
		js:     js,
		taskID: taskID,
		logger: logger,
	}
}

func (w *NatsWriter) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(string(p))
	if msg == "" {
		return len(p), nil
	}
	payload, _ := json.Marshal(spec.AgentResponse{
		Message: msg,
		Status:  spec.StatusPass,
	})
	if _, err := w.js.Publish(w.ctx, fmt.Sprintf(spec.TASK_RESP_SUB, w.taskID), payload); err != nil {
		w.logger.Err(err)
		// Don't return error — let the stream continue even if a single publish fails
	}
	return len(p), nil
}
