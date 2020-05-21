package processor

import (
	"bytes"
	"fmt"
	"net/http"
)

func (p *processor) GetCaddyfile() (content []byte, err error) {
	return p.fileManager.ReadFile(p.caddyfilePath)
}

func (p *processor) SetCaddyfile(content []byte) (err error) {
	r, err := http.NewRequest(http.MethodPost, p.caddyAPIEndpoint, bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	status, respContent, err := p.client.DoHTTPRequest(r)
	if err != nil {
		return err
	} else if status != http.StatusOK {
		return fmt.Errorf("HTTP status code %d", status)
	}
	fmt.Println("Received from API: ", string(respContent))
	return p.fileManager.WriteToFile(p.caddyfilePath, content)
}
