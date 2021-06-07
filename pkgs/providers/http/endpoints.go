package http

func NewEndpoint(method string, path string, handler endpointHandler) *endpoint {
	return &endpoint{
		method:  method,
		path:    path,
		handler: handler,
	}
}
