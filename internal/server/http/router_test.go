package router_test

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/service"
)

func TestNewRouter(t *testing.T) {
	r := router.NewRouter(&service.Services{})

	require.IsType(t, &router.Router{}, r)
}

func TestRouter_Init(t *testing.T) {
	r := router.NewRouter(&service.Services{})
	h := r.Init()
	ts := httptest.NewServer(h)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}
