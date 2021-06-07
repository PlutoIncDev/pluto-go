package httpserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPServer struct {
	router *mux.Router
	server *http.Server
}

func NewHTTPServer(port string) *HTTPServer {
	router := mux.NewRouter()

	return &HTTPServer{
		router: router,
		server: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf("127.0.0.1:%s", port), // todo: port

			// todo: TIMEOUTS
		},
	}
}

func (h *HTTPServer) RegisterHandlerFunc(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	h.router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		handler(writer, request)
	}).Methods(method)
}

func (h *HTTPServer) Run() error {
	return h.server.ListenAndServe()
}
