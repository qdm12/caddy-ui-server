package params

import (
	"github.com/qdm12/golibs/logging"
	libparams "github.com/qdm12/golibs/params"
)

type Reader interface {
	GetListeningPort() (listeningPort, warning string, err error)
	GetLoggerConfig() (encoding logging.Encoding, level logging.Level, nodeID int, err error)
	GetRootURL(setters ...libparams.GetEnvSetter) (rootURL string, err error)
	GetCaddyAPIEndpoint() (endpoint string, err error)
	GetCaddyfilePath() (path string, err error)

	// Version getters
	GetVersion() string
	GetBuildDate() string
	GetVcsRef() string
}

type reader struct {
	envParams libparams.EnvParams
}

func NewReader() Reader {
	return &reader{
		envParams: libparams.NewEnvParams(),
	}
}

func (r *reader) GetListeningPort() (listeningPort, warning string, err error) {
	return r.envParams.GetListeningPort()
}

func (r *reader) GetLoggerConfig() (encoding logging.Encoding, level logging.Level, nodeID int, err error) {
	return r.envParams.GetLoggerConfig()
}

func (r *reader) GetRootURL(setters ...libparams.GetEnvSetter) (rootURL string, err error) {
	return r.envParams.GetRootURL()
}

func (r *reader) GetCaddyAPIEndpoint() (endpoint string, err error) {
	return r.envParams.GetEnv("CADDY_API_ENDPOINT", libparams.Default("http://localhost:2019"))
}

func (r *reader) GetCaddyfilePath() (path string, err error) {
	return r.envParams.GetEnv("CADDYFILE_PATH", libparams.Default("Caddyfile"))
}

func (r *reader) GetVersion() string {
	version, _ := r.envParams.GetEnv("VERSION", libparams.Default("?"), libparams.CaseSensitiveValue())
	return version
}

func (r *reader) GetBuildDate() string {
	buildDate, _ := r.envParams.GetEnv("BUILD_DATE", libparams.Default("?"), libparams.CaseSensitiveValue())
	return buildDate
}

func (r *reader) GetVcsRef() string {
	buildDate, _ := r.envParams.GetEnv("VCS_REF", libparams.Default("?"), libparams.CaseSensitiveValue())
	return buildDate
}
