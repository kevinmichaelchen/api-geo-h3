package main

import (
	"context"
	"github.com/kevinmichaelchen/api-geo-h3/internal/h3"
	"github.com/kevinmichaelchen/api-geo-h3/internal/idl/coop/drivers/h3/v1beta1"
	"github.com/kr/pretty"
	h3Go "github.com/uber/h3-go"
	"log"
)

const (
	lat = 37.775938728915946
	lon = -122.41795063018799
	res = 9
)

func main() {
	ctx := context.Background()

	// Get index from lat/lng
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
	log.Println("Index")
	pretty.Println(index)

	// Get neighboring cells
	kringRes, err := h3.KRing(ctx, &v1beta1.KRingRequest{
		OriginIndex: index,
		K:           1,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Neighbors")
	pretty.Println(kringRes.GetIndexes())

	// K-Ring Distances
	log.Println("KRingDistances")
	kringDistances := h3Go.KRingDistances(h3Go.H3Index(index), 1)
	for _, kr := range kringDistances {
		pretty.Println(h3.FromIndexes(kr))
	}

	// Hex Ring
	log.Println("HexRing")
	hexRing, err := h3Go.HexRing(h3Go.H3Index(index), 1)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println(h3.FromIndexes(hexRing))

	// Get parent
	parentRes, err := h3.H3ToParent(ctx, &v1beta1.H3ToParentRequest{
		Index:      index,
		Resolution: res - 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Parent index = %v\n", parentRes.GetIndex())

	// Get children
	childrenRes, err := h3.H3ToChildren(ctx, &v1beta1.H3ToChildrenRequest{
		ParentIndex: parentRes.GetIndex(),
		Resolution:  res,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Children = %v\n", childrenRes.GetChildrenIndexes())
}
