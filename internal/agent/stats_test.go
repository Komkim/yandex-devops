package agent

import (
	"testing"
	myclient "yandex-devops/provider"
)

func BenchmarkGenerateHas(b *testing.B) {
	v := 186392.0
	key := "123"
	m := []myclient.Metrics{
		{
			ID:    "Alloc",
			MType: "gauge",
			Delta: nil,
			Value: &v,
			Hash:  "",
		},
		{
			ID:    "BuckHashSys",
			MType: "gauge",
			Delta: nil,
			Value: &v,
			Hash:  "",
		},
		{
			ID:    "Frees",
			MType: "gauge",
			Delta: nil,
			Value: &v,
			Hash:  "",
		},
		{
			ID:    "GCCPUFraction",
			MType: "gauge",
			Delta: nil,
			Value: &v,
			Hash:  "",
		},
		{
			ID:    "GCSys",
			MType: "gauge",
			Delta: nil,
			Value: &v,
			Hash:  "",
		},
	}
	for i := 0; i < b.N; i++ {
		generateHas(key, m)
	}
}
