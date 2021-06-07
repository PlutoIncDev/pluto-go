package http

type endpointHandler func()

type endpoint struct {
	method string
	path string
	handler endpointHandler
}

func NewEndpoint(method string, path string, handler endpointHandler) *endpoint {
	return &endpoint{
		method:  method,
		path:    path,
		handler: handler,
	}
}

func (e *endpoint) GetPath() string {
	return e.path
}

func (e *endpoint) GetMethod() string {
	return e.method
}