package router_test

import (
	"log"
	"net/http"
	"net/url"
	"testing"
	"yandex-devops/config"
	router "yandex-devops/internal/server/http"
	"yandex-devops/internal/server/server"
	"yandex-devops/internal/server/service"
	"yandex-devops/storage/memory"

	"github.com/stretchr/testify/require"
)

//func TestNewRouter(t *testing.T) {
//	r := router.NewRouter(&config.Server{}, &service.Services{})
//
//	require.IsType(t, &router.Router{}, r)
//}

func TestRouter_Init(t *testing.T) {
	req := require.New(t)
	s := memory.NewMemStorage()
	cfg, err := config.InitFlagServer()
	req.NoError(err)
	r := router.NewRouter(&cfg.Server, service.NewServices(s))
	srv := server.NewServer(&cfg.HTTP, r.Init())
	go srv.Start()

	log.Println(cfg)

	u := &url.URL{
		Scheme: "http",
		Host:   cfg.HTTP.Address,
	}
	u = u.JoinPath("ping")

	res, err := http.Get(u.String())
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}
