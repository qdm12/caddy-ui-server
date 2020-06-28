package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/qdm12/caddy-ui-server/internal/processor"
	"github.com/qdm12/golibs/errors"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/network"
	"github.com/qdm12/golibs/server"
)

type handler struct {
	logger  logging.Logger
	proc    processor.Processor
	readAll func(r io.Reader) ([]byte, error)
}

func NewHandler(rootURL string, proc processor.Processor, logger logging.Logger) http.HandlerFunc {
	h := &handler{
		proc:    proc,
		logger:  logger,
		readAll: ioutil.ReadAll,
	}
	ipManager := network.NewIPManager(logger)
	return func(w http.ResponseWriter, r *http.Request) {
		ip, err := ipManager.GetClientIP(r)
		if err != nil {
			logger.Error(err)
		}
		logger.Info("HTTP %s %s from %s", r.Method, r.URL.Path, ip)
		path := strings.TrimPrefix(r.URL.Path, rootURL)
		switch {
		case r.Method == http.MethodGet && !strings.HasPrefix(path, "/api"):
			http.ServeFile(w, r, "./ui/"+path)
		case r.Method == http.MethodGet && path == "/api/config":
			h.getCaddyConfig(w)
		case r.Method == http.MethodPost && path == "/api/load":
			h.setCaddyConfig(w, r)
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
