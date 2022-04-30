package h3

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"github.com/uber/h3-go"
)

func KRing(ctx context.Context, r *v1beta1.KRingRequest) (*v1beta1.KRingResponse, error) {
	indexes := h3.KRing(h3.H3Index(r.GetOriginIndex()), int(r.GetK()))
	var out []uint64
	for _, index := range indexes {
		out = append(out, uint64(index))
	}
	return &v1beta1.KRingResponse{Indexes: out}, nil
}
