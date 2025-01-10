package handler

import (
	"net/http"
	"strconv"

	"github.com/Durga-chikkala/unique-request-counter/model"
	"github.com/Durga-chikkala/unique-request-counter/service"
)

type Handler struct {
	service.UniqueRequestCounter
}

func New(svc service.UniqueRequestCounter) *Handler {
	return &Handler{
		UniqueRequestCounter: svc,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed"))

		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed"))

		return
	}

	endpoint := r.URL.Query().Get("endpoint")

	f := model.Filter{Id: int64(idInt), Endpoint: endpoint}

	err = h.UniqueRequestCounter.Get(r.Context(), f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
