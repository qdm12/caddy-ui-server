package processor

import (
	"time"

	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/network"
)

// Processor has methods to process data and return results
type Processor interface {
	GetCaddyConfig() (jsonContent []byte, err error)
	SetCaddyConfig(jsonContent []byte) (err error)
}

type processor struct {
	caddyAPIEndpoint string
	client           network.Client
	logger           logging.Logger
}

// NewProcessor creates a new processor object
func NewProcessor(caddyAPIEndpoint string, logger logging.Logger) Processor {
	return &processor{
		caddyAPIEndpoint: caddyAPIEndpoint,
		client:           network.NewClient(time.Second),
		logger:           logger,
	}
}
