package router

import (
	"compress/gzip"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"yandex-devops/storage"
)

func (h *Router) SaveOrUpdate(c *gin.Context) {
	var mtr storage.Metrics

	if err := json.NewDecoder(c.Request.Body).Decode(&mtr); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if checkHas, err := h.services.StorageService.CheckHash(mtr, h.cfg.Key); err != nil || !checkHas {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := h.services.StorageService.SaveOrUpdateOne(mtr, h.cfg.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, r)
}

// Deprecated: Old version api
func (h *Router) SaveOrUpdateOld(c *gin.Context) {

	t := c.Param("t")
	n := c.Param("n")
	v := c.Param("v")

	m := storage.Metrics{MType: t, ID: n}

	switch t {
	case "counter":
		vf, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad value")
			return
		}
		m.Delta = &vf

	case "gauge":
		vc, err := strconv.ParseFloat(v, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad value")
			return
		}
		m.Value = &vc
	default:
		c.JSON(http.StatusNotImplemented, "Bad value type!")
		return
	}

	r, err := h.services.StorageService.SaveOrUpdateOne(m, h.cfg.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, r)
}

func (h *Router) GetByKey(c *gin.Context) {
	var mtr storage.Metrics

	if err := json.NewDecoder(c.Request.Body).Decode(&mtr); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	str, err := h.services.StorageService.GetByKey(mtr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if str == nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}
	//
	//if len(h.cfg.Key) >= 0 {
	//	mtr.Hash = hex.EncodeToString(h.services.StorageService.GenerageHash(mtr, h.cfg.Key))
	//}

	log.Println(str)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, str)
}

// Deprecated: Old version Api
func (h *Router) GetByKeyOld(c *gin.Context) {

	n := c.Param("n")
	t := c.Param("t")

	mm, err := h.services.StorageService.GetByKey(storage.Metrics{ID: n})
	if err != nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}

	if mm == nil {
		c.JSON(http.StatusNotFound, mm)
		return
	}

	if mm.MType != t {
		c.JSON(http.StatusNotFound, "Bad type")
		return
	}

	switch t {
	case "gauge":
		c.JSON(http.StatusOK, mm.Value)
		return
	case "counter":
		c.JSON(http.StatusOK, mm.Delta)
		return
	default:
		c.JSON(http.StatusBadRequest, err)
		return
	}
}

func (h *Router) GetAll(c *gin.Context) {

	mm, err := h.services.StorageService.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, "Bad key")
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.JSON(http.StatusOK, mm)
}

func (h *Router) Ping(c *gin.Context) {

	ctx := context.Background()
	db, err := sql.Open("pgx", h.cfg.DatabaseDSN)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error connect database")
		return
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error connect database")
		return
	}

	c.JSON(http.StatusOK, "Pong")
}

func (h *Router) SetAll(c *gin.Context) {
	var metrics []storage.Metrics

	if err := json.NewDecoder(c.Request.Body).Decode(&metrics); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for _, m := range metrics {
		if checkHas, err := h.services.StorageService.CheckHash(m, h.cfg.Key); err != nil || !checkHas {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	}
	r, err := h.services.StorageService.SaveOrUpdateAll(metrics, h.cfg.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(c.Writer).Encode(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "ok")

}

func (h *Router) gzipMiddleware(c *gin.Context) {
	if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "qzip") {
		c.Next()
		return
	}
	gz, err := gzip.NewWriterLevel(c.Writer, gzip.BestSpeed)
	if err != nil {
		io.WriteString(c.Writer, err.Error())
		return
	}
	defer gz.Close()

	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Next()
}
