package h3

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"github.com/uber/h3-go"
)

func H3ToParent(ctx context.Context, r *v1beta1.H3ToParentRequest) (*v1beta1.H3ToParentResponse, error) {
	parentIndex := h3.ToParent(h3.H3Index(r.GetIndex()), int(r.GetResolution()))
	return &v1beta1.H3ToParentResponse{Index: fromIndex(parentIndex)}, nil
}
