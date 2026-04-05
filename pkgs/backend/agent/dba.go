package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/jsonstream"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/moby/moby/client"
	"github.com/nats-io/nats.go/jetstream"
)

func (a *Agent) CreateDatabase(ctx context.Context, task_id string, msg jetstream.Msg) error {
	var dbConfig spec.CreateDB
	err := json.Unmarshal(msg.Data(), &dbConfig)
	if err != nil {
		a.logger.Err(err)
		msg.TermWithReason("Unable to convert payload to required format")
		return err
	}
	mesg, _ := json.Marshal(spec.AgentResponse{
		Message: fmt.Sprintf("Pulling Image %s", dbConfig.Image),
	})
	a.logger.Debug().Msgf("Pulling Image %s", dbConfig.Image)

	_, err = a.js.Publish(ctx, fmt.Sprintf(spec.TASK_RESP_SUB, task_id), mesg)
	if err != nil {
		return err
	}

	imgPullResp, err := a.dc.ImagePull(ctx, dbConfig.Image, client.ImagePullOptions{})
	if err != nil {
		errM := msg.TermWithReason(fmt.Sprintf("Error pulling image %s: %v", dbConfig.Image, err.Error()))
		if errM != nil {
			a.logger.Err(errM)
			// return err
		}
		mesg, _ := json.Marshal(spec.AgentResponse{
			Message: fmt.Sprintf("Error pulling image %s: %v", dbConfig.Image, err.Error()),
			Status:  spec.StatusFail,
		})
		_, errP := a.js.Publish(ctx, fmt.Sprintf(spec.TASK_RESP_SUB, task_id), mesg)
		if errP != nil {
			a.logger.Err(errP)
			// return err
		}
		return err
	}

	natsWriter := jsonstream.NewNatsWriter(ctx, a.js, task_id, a.logger)
	if err := jsonstream.Display(ctx, imgPullResp, natsWriter); err != nil {
		errMsg := fmt.Sprintf("Error streaming image pull for %s: %v", dbConfig.Image, err)
		if termErr := msg.TermWithReason(errMsg); termErr != nil {
			a.logger.Err(termErr)
			msg, _ := json.Marshal(spec.AgentResponse{
				Message: fmt.Sprintf("Error pulling image %s: %v", dbConfig.Image, err.Error()),
				Status:  spec.StatusFail,
			})
			a.js.Publish(ctx, fmt.Sprintf(spec.TASK_RESP_SUB, task_id), msg)
		}
		return err
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
	case "create_db":
		err := a.CreateDatabase(ctx, taskID, msg)
		return err

	case "delete_db":
		var db spec.DeleteDB
		err := json.Unmarshal(msg.Data(), &db)
		if err != nil {
			msg.TermWithReason("Unable to convert payload to required format")
			return err
		}
		return a.DeleteDatabase(taskID, db)

	default:
		return fmt.Errorf("unknown task type: %s", taskType)
	}
}
