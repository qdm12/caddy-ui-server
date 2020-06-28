package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qdm12/golibs/errors"
)

func (p *processor) GetCaddyConfig() (jsonContent []byte, err error) {
	r, err := http.NewRequest(http.MethodGet, p.caddyAPIEndpoint+"/config", nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	status, jsonContent, err := p.client.DoHTTPRequest(r)
	if err != nil {
		return nil, err
	}
	p.logger.Info("Caddy (get config) responded HTTP status %d with content: %s", status, string(jsonContent))
	if status != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code %d", status)
	}
	return jsonContent, nil
}

func (p *processor) SetCaddyConfig(jsonContent []byte) (err error) {
	r, err := http.NewRequest(http.MethodPost, p.caddyAPIEndpoint+"/load", bytes.NewBuffer(jsonContent))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	status, jsonContent, err := p.client.DoHTTPRequest(r)
	if err != nil {
		return err
	}
	p.logger.Info("Caddy (set config) responded HTTP status %d with content: %s", status, string(jsonContent))
	if status == http.StatusOK {
		return nil
	}
	response := struct {
		Error string `json:"error"`
	}{}
	if err := json.Unmarshal(jsonContent, &response); err != nil {
		return err
	}
	switch status {
	case http.StatusBadRequest:
		err = errors.NewBadRequest(response.Error)
	case http.StatusConflict:
		err = errors.NewConflict(response.Error)
	}
	return err
}
