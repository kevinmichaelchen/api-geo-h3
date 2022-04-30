package h3

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func H3Distance(ctx context.Context, r *v1beta1.H3DistanceRequest) (*v1beta1.H3DistanceResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
