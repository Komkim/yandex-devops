package router_test

import (
	router "Komkim/go-musthave-devops-tpl/cmd/server/internal/http"
	"Komkim/go-musthave-devops-tpl/cmd/server/internal/service"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
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
