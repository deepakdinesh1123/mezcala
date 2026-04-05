package api

import (
	"context"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
)

func (s *Server) CreateSyncRule(ctx context.Context, req spec.CreateSyncRuleRequestObject) (spec.CreateSyncRuleResponseObject, error) {
	return spec.CreateSyncRule201JSONResponse{}, nil
}

func (s *Server) DeleteSyncRule(ctx context.Context, req spec.DeleteSyncRuleRequestObject) (spec.DeleteSyncRuleResponseObject, error) {
	return spec.DeleteSyncRule204Response{}, nil
}

func (s *Server) ExecuteSyncRule(ctx context.Context, req spec.ExecuteSyncRuleRequestObject) (spec.ExecuteSyncRuleResponseObject, error) {
	return spec.ExecuteSyncRule202JSONResponse{}, nil
}

func (s *Server) UpdateSyncRule(ctx context.Context, req spec.UpdateSyncRuleRequestObject) (spec.UpdateSyncRuleResponseObject, error) {
	return spec.UpdateSyncRule200JSONResponse{}, nil
}
