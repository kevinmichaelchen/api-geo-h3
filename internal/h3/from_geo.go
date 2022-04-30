package h3

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"github.com/uber/h3-go"
)

func FromGeo(ctx context.Context, r *v1beta1.GeoToH3Request) (*v1beta1.GeoToH3Response, error) {
	index := h3.FromGeo(h3.GeoCoord{
		Latitude:  r.GetCoordinates().GetLatitude(),
		Longitude: r.GetCoordinates().GetLongitude(),
	}, int(r.GetResolution()))
	return &v1beta1.GeoToH3Response{Index: uint64(index)}, nil
}
