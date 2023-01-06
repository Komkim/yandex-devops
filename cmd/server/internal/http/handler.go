package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/storage"
	"strconv"
)

func (h *Router) SaveOrUpdate(c *gin.Context) {

	t := c.Param("t")
	n := c.Param("n")
	v := c.Param("v")

	if v == "" {
		c.JSON(http.StatusInternalServerError, "Bad value")
	}

	var m storage.Metric

	switch t {
	case "counter":
		mm, err := h.services.MemStorage.GetByKey(n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		cc, _ := strconv.Atoi(mm.Value)
		cv, _ := strconv.Atoi(v)

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
