package http

func NewProvider(port string) *Provider {
	return &Provider{
		// config
		port: port,

		// internal
		finished: false,
	}
}

func (p *Provider) RegisterEndpoint(method string, path string, handler endpointHandler) {
	e := NewEndpoint(method, path, handler)
	p.endpoints = append(p.endpoints, e)
}
