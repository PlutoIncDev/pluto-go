package rpc

func NewProvider() *Provider {
	return &Provider{
		// config

		// internal
		finished: false,
	}
}
/*
TODO: PROTOCOL BUFFERS FOR SENDING DATA?
 */
// EXAMPLE FROM HTTP SERVER VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV
//func (p *Provider) RegisterEndpoint(method string, path string, handler endpointHandler) {
//	e := NewEndpoint(method, path, handler)
//	p.endpoints = append(p.endpoints, e)
//}
