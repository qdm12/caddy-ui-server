package handlers

import (
	"net/http"
	"strings"

	"github.com/qdm12/caddy-ui-server/internal/processor"
	"github.com/qdm12/golibs/errors"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/server"
)

type handler struct {
	logger logging.Logger
	proc   processor.Processor
}

func NewHandler(rootURL string, proc processor.Processor, logger logging.Logger) http.HandlerFunc {
	h := &handler{
		proc:   proc,
		logger: logger,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, rootURL)
		switch {
		case r.Method == http.MethodGet && path == "/caddyfile":
			h.getCaddyfile(w, r)
		case r.Method == http.MethodPut && path == "/caddyfile":
			h.setCaddyfile(w, r)
		default:
			h.respondError(w, errors.NewBadRequest("invalid %s request at %s", r.Method, path))
		}
	}
}

func (h *handler) respondWrapper(w http.ResponseWriter, setters ...server.ResponseSetter) {
	err := server.Respond(w, setters...)
	if err != nil {
		h.logger.Warn("cannot respond to client: %s", err)
	}
}

func (h *handler) respondError(w http.ResponseWriter, err error) {
	result := struct {
		Error string `json:"error"`
	}{"null"}
	if err != nil {
		result.Error = err.Error()
	}
	status := errors.HTTPStatus(err)
	h.respondWrapper(w, server.Status(status), server.JSON(result))
}
