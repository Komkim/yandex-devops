package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yandex-devops/storage"
)

func (h *Router) SaveOrUpdate(c *gin.Context) {
	var mtr storage.Metrics

	if err := json.NewDecoder(c.Request.Body).Decode(&mtr); err != nil {
		c.JSON(http.StatusBadRequest, "Bad value")
		return
	}

	if err := h.services.Mss.SaveOrUpdateOne(mtr); err != nil {
		c.JSON(http.StatusBadRequest, "Bad value")
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

// Deprecated: Old version api
func (h *Router) SaveOrUpdateOld(c *gin.Context) {

	t := c.Param("t")
	n := c.Param("n")
	v := c.Param("v")

	if _, err := strconv.ParseFloat(v, 64); err != nil {
		c.JSON(http.StatusBadRequest, "Bad value")
		return
	}

	var m storage.Metrics
	var val float64

	switch t {
	case "counter":

		//m := storage.Metrics{
		//ID: n,
		//MType: t,
		//Value: v
		//}

		var cc float64
		m, err := h.services.Mss.GetByKey(storage.Metrics{ID: n})
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		if m == (storage.Metrics{}) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		//if m != (storage.Metrics{}) {
		//	cc, err = strconv.ParseFloat(m.Value, 64)
		//	if err != nil {
		//		c.JSON(http.StatusBadRequest, "Bad value")
		//		return
		//	}
		//}
		if m.Value != nil {
			cc = *m.Value
		}

		cv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad value")
			return
		}

		val = cc + cv

	case "gauge":
		vc, err := strconv.ParseFloat(v, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad value")
			return
		}
		val = vc
	default:
		c.JSON(http.StatusNotImplemented, "Bad value type!")
		return
	}

	m.MType = t
	m.ID = n
	m.Value = &val

	err := h.services.Mss.SaveOrUpdateOne(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "Ok")
}

func (h *Router) GetByKey(c *gin.Context) {
	var mtr storage.Metrics

	if err := json.NewDecoder(c.Request.Body).Decode(&mtr); err != nil {
		c.JSON(http.StatusBadRequest, "Bad value")
		return
	}

	if str, err := h.services.Mss.GetByKey(mtr); err != nil {
		c.JSON(http.StatusBadRequest, "Bad value")
		return
	} else if str == (storage.Metrics{}) {
		c.JSON(http.StatusNotFound, err)
		return
	} else {
		c.JSON(http.StatusOK, str)
	}
}

// Deprecated: Old version Api
func (h *Router) GetByKeyOld(c *gin.Context) {

	n := c.Param("n")
	t := c.Param("t")

	mm, err := h.services.Mss.GetByKey(storage.Metrics{ID: n})
	if err != nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}

	if mm == (storage.Metrics{}) {
		c.JSON(http.StatusNotFound, mm)
		return
	}

	if mm.MType != t && mm != (storage.Metrics{}) {
		c.JSON(http.StatusNotFound, "Bad type")
		return
	}

	switch t {
	case "gauge":
		//if r, err := strconv.ParseFloat(mm.Value, 64); err != nil {
		//	c.JSON(http.StatusBadRequest, err)
		//	return
		//} else {
		//	c.JSON(http.StatusOK, r)
		//	return
		//}
		c.JSON(http.StatusOK, mm.Value)
		return
	case "counter":
		//if r, err := strconv.Atoi(mm.Value); err != nil {
		//	c.JSON(http.StatusBadRequest, err)
		//	return
		//} else {
		//	c.JSON(http.StatusOK, r)
		//	return
		//}
		c.JSON(http.StatusOK, mm.Delta)
		return
	default:
		c.JSON(http.StatusBadRequest, err)
		return
	}
}

func (h *Router) GetAll(c *gin.Context) {

	mm, err := h.services.Mss.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}

	c.JSON(http.StatusOK, mm)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
