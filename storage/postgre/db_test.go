package postgresql

//
//import (
//	"context"
//	"github.com/stretchr/testify/require"
//	"yandex-devops/storage"
//	"yandex-devops/test/fixtures"
//
//	"testing"
//)
//
//func TestPostgreStorage_SetOne(t *testing.T) {
//	if testing.Short() {
//		t.Skip(skipTestMessage)
//	}
//
//	var (
//		req   = require.New(t)
//		ctx   = context.Background()
//		value = float64(3702784)
//		delta = int64(4)
//		repo  = getStorageRepo(req)
//	)
//
//	fixtures.ExecuteFixture(ctx, repo.PGXpool, fixtures.CleanupFixture{})
//
//	metric := []storage.Metrics{
//		{ID: "HeapSys", MType: "gauge", Value: &value, Hash: "e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},
//		{ID: "PollCount", MType: "counter", Delta: &delta, Hash: "f2b586d3bd28e23820ebd4c0149104791821d5e25b96b2d241c383fd1e5b0668"},
//	}
//
//	test := func(m storage.Metrics) func(t *testing.T) {
//		return func(t *testing.T) {
//			_, err := repo.SetOne(m)
//			req.NoError(err)
//
//			r, err := repo.GetOne(m.ID)
//			req.NoError(err)
//
//			req.Equal(m.ID, r.ID)
//			req.Equal(m.MType, r.MType)
//		}
//	}
//
//	t.Run("insert type gauge", test(metric[0]))
//	t.Run("insert type counter", test(metric[1]))
//}
//
//func TestPostgreStorage_GetOne(t *testing.T) {
//	if testing.Short() {
//		t.Skip(skipTestMessage)
//	}
//
//	var (
//		req   = require.New(t)
//		ctx   = context.Background()
//		value = float64(3702784)
//		delta = int64(4)
//		repo  = getStorageRepo(req)
//	)
//
//	fixtures.ExecuteFixture(ctx, repo.PGXpool, fixtures.CleanupFixture{})
//
//	metrics := []storage.Metrics{
//		{ID: "HeapSys", MType: "gauge", Value: &value, Hash: "e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},
//		{ID: "PollCount", MType: "counter", Delta: &delta, Hash: "f2b586d3bd28e23820ebd4c0149104791821d5e25b96b2d241c383fd1e5b0668"},
//	}
//
//	for _, m := range metrics {
//		_, err := repo.SetOne(m)
//		req.NoError(err)
//	}
//
//	test := func(metric *storage.Metrics, isError bool) func(t *testing.T) {
//		return func(t *testing.T) {
//			actual, err := repo.GetOne(metric.ID)
//			if isError {
//				req.Error(err)
//			} else {
//				req.NoError(err)
//			}
//			req.Equal(metric.ID, actual.ID)
//			req.Equal(metric.MType, actual.MType)
//			req.Equal(metric.Delta, actual.Delta)
//			req.Equal(metric.Value, actual.Value)
//			req.Equal(metric.Hash, actual.Hash)
//		}
//	}
//
//	t.Run("select type gauge", test(&metrics[0], false))
//	t.Run("select type counter", test(&metrics[1], false))
//}
//
//func TestPostgreStorage_GetAll(t *testing.T) {
//	if testing.Short() {
//		t.Skip(skipTestMessage)
//	}
//
//	var (
//		req   = require.New(t)
//		ctx   = context.Background()
//		value = float64(3702784)
//		delta = int64(4)
//		repo  = getStorageRepo(req)
//	)
//
//	fixtures.ExecuteFixture(ctx, repo.PGXpool, fixtures.CleanupFixture{})
//
//	metrics := []storage.Metrics{
//		{ID: "HeapSys", MType: "gauge", Value: &value, Hash: "e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},
//		{ID: "PollCount", MType: "counter", Delta: &delta, Hash: "f2b586d3bd28e23820ebd4c0149104791821d5e25b96b2d241c383fd1e5b0668"},
//	}
//	_, err := repo.SetAll(metrics)
//	req.NoError(err)
//
//	test := func(isError bool) func(t *testing.T) {
//		return func(t *testing.T) {
//			actual, err := repo.GetAll()
//			if isError {
//				req.Error(err)
//			} else {
//				req.NoError(err)
//			}
//
//			req.Equal(metrics, actual)
//		}
//	}
//
//	t.Run("select all", test(false))
//}
//
//func TestPostgreStorage_SetAll(t *testing.T) {
//	if testing.Short() {
//		t.Skip(skipTestMessage)
//	}
//
//	var (
//		req   = require.New(t)
//		ctx   = context.Background()
//		value = float64(3702784)
//		delta = int64(4)
//		repo  = getStorageRepo(req)
//	)
//
//	fixtures.ExecuteFixture(ctx, repo.PGXpool, fixtures.CleanupFixture{})
//
//	metrics := []storage.Metrics{
//		{ID: "PollCount", MType: "counter", Delta: &delta, Hash: "f2b586d3bd28e23820ebd4c0149104791821d5e25b96b2d241c383fd1e5b0668"},
//		{ID: "HeapSys", MType: "gauge", Value: &value, Hash: "e3ec1cae0b022f109fada933959833ee75a54c58900c6fe6eca8d70195df13e5"},
//	}
//
//	test := func(m []storage.Metrics) func(t *testing.T) {
//		return func(t *testing.T) {
//			_, err := repo.SetAll(m)
//			req.NoError(err)
//
//			r, err := repo.GetAll()
//			req.NoError(err)
//
//			req.Equal(m, r)
//		}
//	}
//
//	t.Run("select all", test(metrics))
//}
