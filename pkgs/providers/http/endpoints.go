package http

type endpointHandler func()

type endpoint struct {
	method Method
	path string
	handler endpointHandler
}

func NewEndpoint(method Method, path string, handler endpointHandler) *endpoint {
	return &endpoint{
		method:  method,
		path:    path,
		handler: handler,
	}
}

func (e *endpoint) GetPath() string {
	return e.path
}

func (e *endpoint) GetMethod() Method {
	return e.method
}