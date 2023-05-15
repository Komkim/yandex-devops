package agent

import (
	"github.com/stretchr/testify/require"
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

func TestGenerateHas(t *testing.T) {
	req := require.New(t)
	v1 := float64(25109962752)
	v2 := float64(9984978944)
	metrics := []myclient.Metrics{
		{ID: "TotalMemory", MType: "gauge", Delta: nil, Value: &v1, Hash: ""},
		{ID: "FreeMemory", MType: "gauge", Delta: nil, Value: &v2, Hash: ""},
	}

	resultHas := map[string]string{
		"TotalMemory": "b9e085ecd0de555985bf2fb856750e8d222bc6a0b9788682b0c73e533d70e5f8",
		"FreeMemory":  "2402b4866e6f083757e34801eef7f2638d345eb763b0af6558f181bf1fbad7cf",
	}
	key := "123"

	test := func(key string, m []myclient.Metrics) func(t *testing.T) {
		return func(t *testing.T) {
			result := generateHas(key, m)

			for _, v := range result {
				//r, ok := resultHas[v.ID]
				//if ok {
				req.Equal(resultHas[v.ID], v.Hash)
				//}
			}
		}
	}

	t.Run("test generate has", test(key, metrics))
}
