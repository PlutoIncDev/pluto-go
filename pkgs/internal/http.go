package internal

type HttpMethod string

const (
	GetHttpMethod HttpMethod = "GET"
	PostHttpMethod = "POST"
	PatchHttpMethod = "PATCH"
	DeleteHttpMethod = "DELETE"
)

type httpHandlerFunc = func()

type httpHandler struct {
	Method HttpMethod
	Handler httpHandlerFunc
	Path string
}

func newHttpHandler(method HttpMethod, path string, handler httpHandlerFunc) *httpHandler {
	return &httpHandler{
		Method: method,
		Handler: handler,
		Path: path,
	}
}

type Http struct {
	handlers []*httpHandler
}

func newHttp() *Http {
	return new(Http)
}

func (h *Http) registerHttpHandler(method HttpMethod, path string, handler httpHandlerFunc) {
	h.handlers = append(h.handlers, newHttpHandler(method, path, handler))
}

func (h *Http) GetHandlers() []*httpHandler {
	return h.handlers
}

func (h *Http) Get(path string, handler httpHandlerFunc) {
	h.registerHttpHandler(GetHttpMethod, path, handler)
}

func (h *Http) Post(path string, handler httpHandlerFunc) {
	h.registerHttpHandler(PostHttpMethod, path, handler)
}

func (h *Http) Patch(path string, handler httpHandlerFunc) {
	h.registerHttpHandler(PatchHttpMethod, path, handler)
}

func (h *Http) Delete(path string, handler httpHandlerFunc) {
	h.registerHttpHandler(DeleteHttpMethod, path, handler)
}