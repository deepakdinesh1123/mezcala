package agent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/nats-io/nats.go/jetstream"
)

func (a *Agent) CreateDatabase(task_id string, dbConfig spec.DatabaseConfig) error {
	return nil
}

func (a *Agent) DeleteDatabase(task_id string, db spec.DeleteDB) error {
	return nil
}

func (a *Agent) handleAdmin(msg jetstream.Msg) error {
	parts := strings.Split(msg.Subject(), ".")
	if len(parts) < 4 {
		return fmt.Errorf("invalid subject format: %s", msg.Subject())
	}

	taskType := parts[2]
	taskID := parts[3]

	switch taskType {
	case "createDB":
		var dbConfig spec.DatabaseConfig
		err := json.Unmarshal(msg.Data(), &dbConfig)
		if err != nil {
			msg.TermWithReason("Unable to convert payload to required format")
			return err
		}
		a.CreateDatabase(taskID, dbConfig)

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
