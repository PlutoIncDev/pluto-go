package http

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.New(),
	}
}

func (s *Server) AddMiddleware() {

}