package router

import (
	"Komkim/go-musthave-devops-tpl/cmd/server/internal/service"
	"net/http"
)

type Router struct {
	services *service.Services
}

func NewRouter(s *service.Services) *Router {
	return &Router{s}
}

func (r *Router) Init() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/update/", r.SaveOrUpdate)
	mux.HandleFunc("/ping", Ping)

	return mux
}
