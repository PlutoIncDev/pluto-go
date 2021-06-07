package http

import "pluto/pkgs/internal/httpserver"

type endpointHandler func(ctx *Context)

type endpoint struct {
	method Method
	path string
	handler endpointHandler
}

type Provider struct {
	port string
	endpoints []*endpoint
	server *httpserver.HTTPServer
}

type Method string

const (
	GetMethod Method = "GET"
	PostMethod Method = "POST"
	PatchMethod Method = "PATCH"
	DeleteMethod Method = "DELETE"
	OptionsMethod Method = "OPTIONS"
)

type Context struct {
	 
}

type Headers struct {

}