package main

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/h3"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"log"
)

const (
	lat = 37.775938728915946
	lon = -122.41795063018799
	res = 9
)

func main() {
	ctx := context.Background()
	geoRes, err := h3.FromGeo(ctx, &v1beta1.GeoToH3Request{
		Coordinates: &v1beta1.LatLng{
			Latitude:  lat,
			Longitude: lon,
		},
		Resolution: res,
	})
	if err != nil {
		log.Fatal(err)
	}
	index := geoRes.GetIndex()
	log.Println("Index = ", index)

	kringRes, err := h3.KRing(ctx, &v1beta1.KRingRequest{
		OriginIndex: index,
		K:           1,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Neighbors = %v\n", kringRes.GetIndexes())
}
