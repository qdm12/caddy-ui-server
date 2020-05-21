package processor

import (
	"time"

	"github.com/qdm12/golibs/files"
	"github.com/qdm12/golibs/network"
)

// Processor has methods to process data and return results
type Processor interface {
	GetCaddyfile() (content []byte, err error)
	SetCaddyfile(content []byte) (err error)
}

type processor struct {
	caddyAPIEndpoint string
	caddyfilePath    string
	fileManager      files.FileManager
	client           network.Client
}

// NewProcessor creates a new processor object
func NewProcessor(caddyAPIEndpoint, caddyfilePath string, fileManager files.FileManager) Processor {
	return &processor{
		caddyAPIEndpoint: caddyAPIEndpoint,
		caddyfilePath:    caddyfilePath,
		fileManager:      fileManager,
		client:           network.NewClient(time.Second),
	}
}
