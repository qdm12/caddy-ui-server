package handlers

import (
	"net/http"

	"github.com/qdm12/golibs/server"
)

func (h *handler) getCaddyfile(w http.ResponseWriter) {
	content, err := h.proc.GetCaddyfile()
	if err != nil {
		h.respondError(w, err)
		return
	}
	h.respondWrapper(w,
		server.Status(http.StatusOK),
		server.Bytes(content, "text/caddyfile"))
}

func (h *handler) setCaddyfile(w http.ResponseWriter, r *http.Request) {
	content, err := h.readAll(r.Body)
	defer func() {
		if err := r.Body.Close(); err != nil {
			h.logger.Error(err)
		}
	}()
	if err != nil {
		h.respondError(w, err)
		return
	}
	if err := h.proc.SetCaddyfile(content); err != nil {
		h.respondError(w, err)
		return
	}
	h.respondWrapper(w, server.Status(http.StatusOK))
}
