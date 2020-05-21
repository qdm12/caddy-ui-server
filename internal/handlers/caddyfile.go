package handlers

import (
	"net/http"

	"github.com/qdm12/golibs/server"
)

func (h *handler) getCaddyfile(w http.ResponseWriter, r *http.Request) {
	h.respondWrapper(w, server.Status(http.StatusOK))
}

func (h *handler) setCaddyfile(w http.ResponseWriter, r *http.Request) {
	h.respondWrapper(w, server.Status(http.StatusOK))
}
