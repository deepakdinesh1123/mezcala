package api

import (
	"context"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/spec"
)

func (s *Server) CheckHealth(ctx context.Context, req spec.CheckHealthRequestObject) (spec.CheckHealthResponseObject, error) {
	return spec.CheckHealth200JSONResponse{}, nil
}
