package metric

import (
	"net/http"

	"github.com/RndIT/go_demos/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
	Logger logging.Logger
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(204)
}
