package http

import (
	"github.com/gin-gonic/gin"
)

type endpointHandler func(ctx *Context)

type endpoint struct {
	method  Method
	path    string
	handler endpointHandler
}

type Provider struct {
	port      string
	endpoints []*endpoint
	server    *gin.Engine
}

type Method string

const (
	GetMethod     Method = "GET"
	PostMethod    Method = "POST"
	PatchMethod   Method = "PATCH"
	DeleteMethod  Method = "DELETE"
	OptionsMethod Method = "OPTIONS"
)

type Context struct {
	HTTP *gin.Context
}

type H gin.H
