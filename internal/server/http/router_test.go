package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"yandex-devops/config"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/service"

	"github.com/stretchr/testify/require"
)

func TestNewRouter(t *testing.T) {
	r := router.NewRouter(&config.Server{}, &service.Services{})

	require.IsType(t, &router.Router{}, r)
}

func TestRouter_Init(t *testing.T) {
	r := router.NewRouter(&config.Server{}, &service.Services{})
	h := r.Init()
	ts := httptest.NewServer(h)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
