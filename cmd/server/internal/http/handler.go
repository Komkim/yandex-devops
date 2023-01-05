package router

import (
	"Komkim/go-musthave-devops-tpl/cmd/server/storage"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (h *Router) SaveOrUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.URL.Path)
	urlSlice := strings.SplitN(r.URL.Path, "/", 5)
	var m storage.Metric

	switch urlSlice[3] {
	case "counter":
		mm, err := h.services.MemStorage.GetByKey(urlSlice[2])
		cc, err := strconv.Atoi(mm.Value)
		c, err := strconv.Atoi(urlSlice[4])

		if err != nil {
			http.Error(w, "Bad value", http.StatusInternalServerError)
			return
		}

		m = storage.Metric{
			Name:  urlSlice[2],
			Type:  urlSlice[3],
			Value: strconv.Itoa(c + cc),
		}
	case "gauge":
		m = storage.Metric{
			Name:  urlSlice[2],
			Type:  urlSlice[3],
			Value: urlSlice[4],
		}
	default:
		http.Error(w, "Bad value type!", http.StatusInternalServerError)
		return
	}

	h.services.MemStorage.SaveOrUpdate(m)

	fmt.Println(h.services.MemStorage.GetAll())

	w.WriteHeader(http.StatusOK)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
