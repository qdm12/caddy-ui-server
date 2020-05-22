package processor

import (
	"bytes"
	"fmt"
	"net/http"
)

func (p *processor) GetCaddyfile() (content []byte, err error) {
	return p.fileManager.ReadFile(p.dataPath + "/Caddyfile")
}

func (p *processor) SetCaddyfile(content []byte) (err error) {
	r, err := http.NewRequest(http.MethodPost, p.caddyAPIEndpoint+"/load", bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "text/caddyfile")
	status, _, err := p.client.DoHTTPRequest(r)
	if err != nil {
		return err
	} else if status != http.StatusOK {
		return fmt.Errorf("HTTP status code %d", status)
	}
	return p.fileManager.WriteToFile(p.dataPath+"/Caddyfile", content)
}
