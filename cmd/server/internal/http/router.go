package router

import (
	"Komkim/go-musthave-devops-tpl/cmd/server/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	services *service.Services
}

func NewRouter(s *service.Services) *Router {
	return &Router{s}
}

func (r *Router) Init() http.Handler {

	mux := gin.Default()

	//mux.Use(gin.Recovery())

	mux.POST("/update/:t/:n/:v", r.SaveOrUpdate)
	mux.GET("/value/:t/:n", r.GetByKey)
	mux.GET("/", r.GetAll)
	mux.GET("/ping", Ping)

	return mux
}
