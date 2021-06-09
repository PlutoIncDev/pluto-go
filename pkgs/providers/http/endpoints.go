package http

func NewEndpoint(method string, path string, handler Handler) *Endpoint {
	return &Endpoint{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
