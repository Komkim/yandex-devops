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

	if _, err := strconv.Atoi(v); err != nil {
		c.JSON(http.StatusBadRequest, "Bad value")
		return
	}

	var m storage.Metric
	var val string

	switch t {
	case "counter":
		var cc int
		m, err := h.services.MemStorage.GetByKey(n)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if m != (storage.Metric{}) {
			cc, err = strconv.Atoi(m.Value)
			if err != nil {
				c.JSON(http.StatusBadRequest, "Bad value")
				return
			}
		}

		cv, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad value")
			return
		}

		val = strconv.Itoa(cv + cc)

	case "gauge":
		val = v
	default:

		c.JSON(http.StatusNotImplemented, "Bad value type!")
		return
	}

	m.Type = t
	m.Name = n
	m.Value = val

	h.services.MemStorage.SaveOrUpdate(m)

	c.JSON(http.StatusOK, "Ok")
}

func (h *Router) GetByKey(c *gin.Context) {
	n := c.Param("n")
	t := c.Param("t")

	mm, err := h.services.MemStorage.GetByKey(n)
	if err != nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}
	if mm.Type != t && mm != (storage.Metric{}) {
		c.JSON(http.StatusNotFound, "Bad type")
		return
	}

	switch t {
	case "gauge":
		r, _ := strconv.ParseFloat(mm.Value, 64)
		c.JSON(http.StatusOK, r)
		return
	case "counter":
		r, _ := strconv.Atoi(mm.Value)
		c.JSON(http.StatusOK, r)
		return
	}

	c.JSON(http.StatusOK, mm.Value)
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
