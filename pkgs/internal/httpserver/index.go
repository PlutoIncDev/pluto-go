package httpserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPServer struct {
	router *mux.Router
	server *http.Server
}

func NewHTTPServer() *HTTPServer {
	router := mux.NewRouter()

	return &HTTPServer{
		router: router,
		server: &http.Server{
			Handler: router,
			Addr:    "127.0.0.1:8000", // todo: port

			// todo: TIMEOUTS
		},
	}
}

func (h *HTTPServer) RegisterHandlerFunc(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	h.router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		handler(writer, request)
	}).Methods(method)
}

func (h *HTTPServer) Start() error {
	return h.server.ListenAndServe()
}
