package code

import (
	"testing"
)

func BenchmarkLoopMap(b *testing.B) {
	req := &QuotationSetReq{}
	req.EstimateID = "123"
	req.Params = FieldMap{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5", "k6": "v6", "k7": "v7", "k8": "v8"}
	for i := 0; i < b.N; i++ {
		LoopMap(req)
	}
}

func BenchmarkJsonMap(b *testing.B) {
	req := &QuotationSetReq{}
	req.EstimateID = "123"
	req.Params = FieldMap{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5", "k6": "v6", "k7": "v7", "k8": "v8"}
	for i := 0; i < b.N; i++ {
		JsonMap(req)
	}
}
func BenchmarkLoopStruct(b *testing.B) {
	req := &InReq{}
	for i := 0; i < b.N; i++ {
		LoopStruct(req)
	}
}

func BenchmarkJsonStruct(b *testing.B) {
	req := &InReq{}
	for i := 0; i < b.N; i++ {
		JsonStruct(req)
	}
}
