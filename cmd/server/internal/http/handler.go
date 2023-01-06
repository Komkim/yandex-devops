package router

import (
	"Komkim/go-musthave-devops-tpl/cmd/server/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Router) SaveOrUpdate(c *gin.Context) {

	t := c.Param("t")
	n := c.Param("n")
	v := c.Param("v")

	var m storage.Metric

	switch t {
	case "counter":
		mm, err := h.services.MemStorage.GetByKey(n)
		cc, _ := strconv.Atoi(mm.Value)
		cv, _ := strconv.Atoi(v)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "Bad value")
			return
		}

		m = storage.Metric{
			Name:  n,
			Type:  t,
			Value: strconv.Itoa(cv + cc),
		}
	case "gauge":
		m = storage.Metric{
			Name:  n,
			Type:  t,
			Value: v,
		}
	default:
		c.JSON(http.StatusInternalServerError, "Bad value type!")
		return
	}

	h.services.MemStorage.SaveOrUpdate(m)

	c.JSON(http.StatusOK, "Ok")
}

func (h *Router) GetByKey(c *gin.Context) {
	n := c.Param("n")
	t := c.Param("t")

	mm, err := h.services.MemStorage.GetByKey(n)
	if err != nil || mm == (storage.Metric{}) {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}
	if mm.Type != t {
		c.JSON(http.StatusNotFound, "Bad type")
		return
	}

	c.JSON(http.StatusOK, mm)
}

func (h *Router) GetAll(c *gin.Context) {

	mm, err := h.services.MemStorage.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}

	c.JSON(http.StatusOK, mm)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
