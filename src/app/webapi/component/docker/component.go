package docker

import (
	"app/webapi/component"
	"github.com/valyala/fasthttp"
	"time"
)

// New returns a new instance of the endpoint.
func New(bc component.Core) *Endpoint {
	return &Endpoint{
		Core: bc,
	}
}

// Endpoint contains the dependencies.
type Endpoint struct {
	component.Core
}

var (
	// HTTPClient global http client object
	HTTPClient *fasthttp.Client = &fasthttp.Client{
		MaxConnsPerHost: 16384, // MaxConnsPerHost  default is 512, increase to 16384
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
	}
)

// Routes will set up the endpoints.
func (p *Endpoint) Routes(router component.IRouter) {
	router.Post("/v1/docker/info", p.Info)
	router.Post("/v1/docker/containers", p.Containers)
}
