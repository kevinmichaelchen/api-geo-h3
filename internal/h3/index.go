package h3

import "github.com/uber/h3-go"

func fromIndexes(in []h3.H3Index) []uint64 {
	var out []uint64
	for _, index := range in {
		out = append(out, uint64(index))
	}
	return out
}

func fromIndex(in h3.H3Index) uint64 {
	return uint64(in)
}
