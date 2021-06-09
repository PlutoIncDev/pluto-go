package types

import (
	"github.com/gin-gonic/gin"
)

type Handler func(ctx *Context)

type Endpoint struct {
	Method  string
	Path    string
	Handler Handler
}

type Provider struct {
	port     string
	finished bool

	endpoints   []*Endpoint
	middlewares []Handler
	server      *gin.Engine
	ginMode     string
}

type Context struct {
	HTTP *gin.Context
}

type H gin.H
