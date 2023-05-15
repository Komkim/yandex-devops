package agent

import (
	"github.com/stretchr/testify/require"
	"runtime"
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
				req.Equal(resultHas[v.ID], v.Hash)
			}
		}
	}

	t.Run("test generate has", test(key, metrics))
}

func TestConvertRuntimeMemory(t *testing.T) {
	var (
		req     = require.New(t)
		counter = int64(1)
		rand    = float64(0.22908455980527204)
	)
	stats := &runtime.MemStats{
		Alloc:         188336,
		TotalAlloc:    188336,
		Sys:           13220880,
		Lookups:       0,
		Mallocs:       734,
		Frees:         17,
		HeapAlloc:     188336,
		HeapSys:       3702784,
		HeapIdle:      2940928,
		HeapInuse:     761856,
		HeapReleased:  2940928,
		HeapObjects:   717,
		StackInuse:    491520,
		StackSys:      491520,
		MSpanInuse:    39744,
		MSpanSys:      48816,
		MCacheInuse:   9600,
		MCacheSys:     15600,
		BuckHashSys:   3886,
		GCSys:         7969240,
		OtherSys:      989034,
		NextGC:        4194304,
		LastGC:        0,
		PauseTotalNs:  0,
		PauseNs:       [256]uint64{},
		PauseEnd:      [256]uint64{},
		NumGC:         0,
		NumForcedGC:   0,
		GCCPUFraction: 0,
		EnableGC:      true,
		DebugGC:       false,
		BySize: [61]struct {
			Size    uint32
			Mallocs uint64
			Frees   uint64
		}{},
	}

	gaugeResult := []float64{188336, 3886, 17, 0, 7969240, 188336, 2940928, 761856, 717, 2940928,
		3702784, 0, 0, 9600, 15600, 39744, 48816, 734, 4194304, 0,
		0, 989034, 0, 491520, 491520, 13220880, 188336, 0.22908455980527204,
	}
	counterResult := int64(1)

	result := []myclient.Metrics{
		{ID: "Alloc", MType: "gauge", Delta: nil, Value: &gaugeResult[0], Hash: ""},
		{ID: "BuckHashSys", MType: "gauge", Delta: nil, Value: &gaugeResult[1], Hash: ""},
		{ID: "Frees", MType: "gauge", Delta: nil, Value: &gaugeResult[2], Hash: ""},
		{ID: "GCCPUFraction", MType: "gauge", Delta: nil, Value: &gaugeResult[3], Hash: ""},
		{ID: "GCSys", MType: "gauge", Delta: nil, Value: &gaugeResult[4], Hash: ""},
		{ID: "HeapAlloc", MType: "gauge", Delta: nil, Value: &gaugeResult[5], Hash: ""},
		{ID: "HeapIdle", MType: "gauge", Delta: nil, Value: &gaugeResult[6], Hash: ""},
		{ID: "HeapInuse", MType: "gauge", Delta: nil, Value: &gaugeResult[7], Hash: ""},
		{ID: "HeapObjects", MType: "gauge", Delta: nil, Value: &gaugeResult[8], Hash: ""},
		{ID: "HeapReleased", MType: "gauge", Delta: nil, Value: &gaugeResult[9], Hash: ""},
		{ID: "HeapSys", MType: "gauge", Delta: nil, Value: &gaugeResult[10], Hash: ""},
		{ID: "LastGC", MType: "gauge", Delta: nil, Value: &gaugeResult[11], Hash: ""},
		{ID: "Lookups", MType: "gauge", Delta: nil, Value: &gaugeResult[12], Hash: ""},
		{ID: "MCacheInuse", MType: "gauge", Delta: nil, Value: &gaugeResult[13], Hash: ""},
		{ID: "MCacheSys", MType: "gauge", Delta: nil, Value: &gaugeResult[14], Hash: ""},
		{ID: "MSpanInuse", MType: "gauge", Delta: nil, Value: &gaugeResult[15], Hash: ""},
		{ID: "MSpanSys", MType: "gauge", Delta: nil, Value: &gaugeResult[16], Hash: ""},
		{ID: "Mallocs", MType: "gauge", Delta: nil, Value: &gaugeResult[17], Hash: ""},
		{ID: "NextGC", MType: "gauge", Delta: nil, Value: &gaugeResult[18], Hash: ""},
		{ID: "NumForcedGC", MType: "gauge", Delta: nil, Value: &gaugeResult[19], Hash: ""},
		{ID: "NumGC", MType: "gauge", Delta: nil, Value: &gaugeResult[20], Hash: ""},
		{ID: "OtherSys", MType: "gauge", Delta: nil, Value: &gaugeResult[21], Hash: ""},
		{ID: "PauseTotalNs", MType: "gauge", Delta: nil, Value: &gaugeResult[22], Hash: ""},
		{ID: "StackInuse", MType: "gauge", Delta: nil, Value: &gaugeResult[23], Hash: ""},
		{ID: "StackSys", MType: "gauge", Delta: nil, Value: &gaugeResult[24], Hash: ""},
		{ID: "Sys", MType: "gauge", Delta: nil, Value: &gaugeResult[25], Hash: ""},
		{ID: "TotalAlloc", MType: "gauge", Delta: nil, Value: &gaugeResult[26], Hash: ""},
		{ID: "PollCount", MType: "counter", Delta: &counterResult, Value: nil, Hash: ""},
		{ID: "RandomValue", MType: "gauge", Delta: nil, Value: &gaugeResult[27], Hash: ""},
	}

	test := func(stats *runtime.MemStats, counter int64, rand float64) func(t *testing.T) {
		return func(t *testing.T) {
			r := convertRuntimeMemory(stats, counter, rand)

			req.Equal(result, r)
		}
	}
	t.Run("test convert runtime memory", test(stats, counter, rand))
}
