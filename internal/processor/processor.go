package processor

// Processor has methods to process data and return results
type Processor interface {
	GetCaddyfile() (content []byte, err error)
	SetCaddyfile(content []byte) (err error)
}

type processor struct {
	caddyAPIEndpoint string
}

// NewProcessor creates a new processor object
func NewProcessor(caddyAPIEndpoint string) Processor {
	return &processor{
		caddyAPIEndpoint: caddyAPIEndpoint,
	}
}
