package h3

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"github.com/uber/h3-go"
)

func H3ToChildren(ctx context.Context, r *v1beta1.H3ToChildrenRequest) (*v1beta1.H3ToChildrenResponse, error) {
	children := h3.ToChildren(h3.H3Index(r.GetParentIndex()), int(r.GetResolution()))
	return &v1beta1.H3ToChildrenResponse{ChildrenIndexes: FromIndexes(children)}, nil
}
