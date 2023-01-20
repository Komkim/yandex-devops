package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yandex-devops/internal/server/service"
)

type Router struct {
	services *service.Services
}

func NewRouter(s *service.Services) *Router {
	return &Router{s}
}

func (h *Router) Init() http.Handler {

	mux := gin.Default()

	//mux.Use(gin.Recovery())

	mux.POST("/update/:t/:n/:v", h.SaveOrUpdateOld)
	mux.POST("/update/:t/", func(c *gin.Context) { c.JSON(http.StatusNotFound, "Not Found") })
	mux.GET("/value/:t/:n", h.GetByKeyOld)
	mux.GET("/value/:t/", func(c *gin.Context) { c.JSON(http.StatusNotFound, "Not Found") })
	mux.GET("/", h.GetAll)
	mux.GET("/ping", Ping)
	mux.POST("/update/", h.SaveOrUpdate)
	mux.POST("/value/", h.GetByKey)

	return mux
}
