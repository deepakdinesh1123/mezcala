package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/moby/moby/client"
	"github.com/nats-io/nats.go/jetstream"
)

func (a *Agent) CreateDatabase(ctx context.Context, task_id string, dbConfig spec.CreateDB) error {
	msg, _ := json.Marshal(spec.AgentResponse{
		Message: fmt.Sprintf("Pulling Image %s", dbConfig.Image),
	})
	_, err := a.js.Publish(ctx, fmt.Sprintf(spec.CREATE_DB_SUB, task_id), msg)
	if err != nil {
		return err
	}

	imgPullResp, err := a.dc.ImagePull(ctx, dbConfig.Image, client.ImagePullOptions{})
	if err != nil {
		msg, _ := json.Marshal(spec.AgentResponse{
			Message: fmt.Sprintf("Error pulling image %s: %v", dbConfig.Image, err.Error()),
			Status:  spec.StatusFail,
		})
		_, err := a.js.Publish(ctx, fmt.Sprintf(spec.CREATE_DB_SUB, task_id), msg)
		if err != nil {
			return err
		}
		return err
	}

	for resp, err := range imgPullResp.JSONMessages(ctx) {
		if err != nil {
			msg, _ := json.Marshal(spec.AgentResponse{
				Message: fmt.Sprintf("Error pulling image %s: %v", dbConfig.Image, err.Error()),
				Status:  spec.StatusFail,
			})
			a.js.Publish(ctx, fmt.Sprintf(spec.CREATE_DB_SUB, task_id), msg)
		} else {
			fmt.Printf("%s", resp.Stream)
			msg, _ := json.Marshal(spec.AgentResponse{
				Message: resp.Stream,
				Status:  spec.StatusPass,
			})
			a.js.Publish(ctx, fmt.Sprintf(spec.CREATE_DB_SUB, task_id), msg)
		}
	}
	return nil
}

func (a *Agent) DeleteDatabase(task_id string, db spec.DeleteDB) error {
	return nil
}

func (a *Agent) handleAdmin(ctx context.Context, msg jetstream.Msg) error {
	parts := strings.Split(msg.Subject(), ".")
	if len(parts) < 4 {
		return fmt.Errorf("invalid subject format: %s", msg.Subject())
	}

	taskType := parts[2]
	taskID := parts[3]

	a.logger.Debug().Msg(taskType)
	a.logger.Debug().Msg(taskID)
	a.logger.Debug().Msg(string(msg.Data()))

	switch taskType {
	case "createDB":
		var dbConfig spec.CreateDB
		err := json.Unmarshal(msg.Data(), &dbConfig)
		if err != nil {
			msg.TermWithReason("Unable to convert payload to required format")
			return err
		}
		a.CreateDatabase(ctx, taskID, dbConfig)

	case "deleteDB":
		var db spec.DeleteDB
		err := json.Unmarshal(msg.Data(), &db)
		if err != nil {
			msg.TermWithReason("Unable to convert payload to required format")
			return err
		}
		a.DeleteDatabase(taskID, db)

	default:
		return fmt.Errorf("unknown task type: %s", taskType)
	}

	return nil
}
