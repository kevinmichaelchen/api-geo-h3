package h3

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"github.com/uber/h3-go"
)

func GetResolution(ctx context.Context, r *v1beta1.H3GetResolutionRequest) (*v1beta1.H3GetResolutionResponse, error) {
	resolution := h3.Resolution(ToIndex(r.GetIndex()))
	return &v1beta1.H3GetResolutionResponse{Resolution: int32(resolution)}, nil
}
