package processor

import (
	"bytes"
	"fmt"
	"net/http"
)

func (p *processor) GetCaddyConfig() (jsonContent []byte, err error) {
	r, err := http.NewRequest(http.MethodGet, p.caddyAPIEndpoint+"/config", nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	status, content, err := p.client.DoHTTPRequest(r)
	if err != nil {
		return nil, err
	} else if status != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code %d", status)
	}
	return content, nil
}

func (p *processor) SetCaddyConfig(jsonContent []byte) (err error) {
	r, err := http.NewRequest(http.MethodPost, p.caddyAPIEndpoint+"/load", bytes.NewBuffer(jsonContent))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	status, _, err := p.client.DoHTTPRequest(r)
	if err != nil {
		return err
	} else if status != http.StatusOK {
		return fmt.Errorf("HTTP status code %d", status)
	}
	return nil
}
