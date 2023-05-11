package router

import (
	"net/http"
	"yandex-devops/config"
	"yandex-devops/internal/server/service"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Router struct {
	services *service.Services
	cfg      *config.Server
}

func NewRouter(cfg *config.Server, s *service.Services) *Router {
	return &Router{cfg: cfg, services: s}
}

func (h *Router) Init() http.Handler {

	mux := gin.Default()
	pprof.Register(mux, "debug/pprof")

	//mux.Use(gin.Recovery())

	//mux.Use(h.gzipMiddleware)
	mux.Use(gzip.Gzip(gzip.BestSpeed))

	mux.POST("/update/:t/:n/:v", h.SaveOrUpdateOld)
	mux.POST("/update/:t/", func(c *gin.Context) { c.JSON(http.StatusNotFound, "Not Found") })
	mux.GET("/value/:t/:n", h.GetByKeyOld)
	mux.GET("/value/:t/", func(c *gin.Context) { c.JSON(http.StatusNotFound, "Not Found") })
	mux.GET("/", h.GetAll)
	mux.GET("/ping", h.Ping)
	mux.POST("/update/", h.SaveOrUpdate)
	mux.POST("/updates/", h.SetAll)
	mux.POST("/value/", h.GetByKey)

	return mux
}
