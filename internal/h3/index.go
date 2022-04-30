package h3

import "github.com/uber/h3-go"

func FromIndexes(in []h3.H3Index) []uint64 {
	var out []uint64
	for _, index := range in {
		out = append(out, FromIndex(index))
	}
	return out
}

func FromIndex(in h3.H3Index) uint64 {
	return uint64(in)
}

func ToIndexes(in []uint64) []h3.H3Index {
	var out []h3.H3Index
	for _, e := range in {
		out = append(out, ToIndex(e))
	}
	return out
}

func ToIndex(in uint64) h3.H3Index {
	return h3.H3Index(in)
}
