package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
	"github.com/google/uuid"
)

func (s *Server) AddDatabase(ctx context.Context, req spec.AddDatabaseRequestObject) (spec.AddDatabaseResponseObject, error) {
	return spec.AddDatabase201JSONResponse{}, nil
}

func (s *Server) CreateDatabase(ctx context.Context, req spec.CreateDatabaseRequestObject) (spec.CreateDatabaseResponseObject, error) {
	task_id := uuid.New()
	payload, err := json.Marshal(spec.CreateDB{
		Image:    *req.Body.Image,
		Name:     *req.Body.Name,
		Password: req.Body.Password,
		Port:     req.Body.Port,
		SslMode:  *req.Body.SslMode,
		Username: req.Body.Username,
	})
	if err != nil {
		return spec.CreateDatabase400JSONResponse{
			BadRequestJSONResponse: spec.BadRequestJSONResponse{
				Code:    "400",
				Message: "Error parsing the request body",
			},
		}, nil
	}

	s.logger.Debug().Msg(fmt.Sprintf(spec.CREATE_DB_SUB, task_id))
	_, err = s.js.Publish(ctx, fmt.Sprintf(spec.CREATE_DB_SUB, task_id), payload)
	if err != nil {
		s.logger.Error().Msg(err.Error())
		return spec.CreateDatabase500JSONResponse{
			InternalServerErrorJSONResponse: spec.InternalServerErrorJSONResponse{
				Code:    "500",
				Message: err.Error(),
			},
		}, nil
	}
	return spec.CreateDatabase202JSONResponse{
		TaskAcceptedJSONResponse: spec.TaskAcceptedJSONResponse{
			TaskId: task_id,
		},
	}, nil
}

func (s *Server) DeleteDatabase(ctx context.Context, req spec.DeleteDatabaseRequestObject) (spec.DeleteDatabaseResponseObject, error) {
	return spec.DeleteDatabase202JSONResponse{}, nil
}

func (s *Server) GetSupportedEngines(ctx context.Context, req spec.GetSupportedEnginesRequestObject) (spec.GetSupportedEnginesResponseObject, error) {

	result := spec.SupportedEnginesResponse{
		{
			Engine:   "Postgres",
			Versions: []string{"16, 17"},
		},
		{
			Engine:   "MySQL",
			Versions: []string{"9.6, 8.4"},
		},
	}

	return spec.GetSupportedEngines200JSONResponse{SupportedEnginesResponseJSONResponse: result}, nil
}

func (s *Server) GetCreateDatabaseTaskStatus(ctx context.Context, req spec.GetCreateDatabaseTaskStatusRequestObject) (spec.GetCreateDatabaseTaskStatusResponseObject, error) {
	return spec.GetCreateDatabaseTaskStatus200JSONResponse{}, nil
}

func (s *Server) GetDatabaseConnectionUrl(ctx context.Context, req spec.GetDatabaseConnectionUrlRequestObject) (spec.GetDatabaseConnectionUrlResponseObject, error) {
	return spec.GetDatabaseConnectionUrl200JSONResponse{}, nil
}

func (s *Server) GetDeleteDatabaseTaskStatus(ctx context.Context, req spec.GetDeleteDatabaseTaskStatusRequestObject) (spec.GetDeleteDatabaseTaskStatusResponseObject, error) {
	return spec.GetDeleteDatabaseTaskStatus200JSONResponse{}, nil
}
