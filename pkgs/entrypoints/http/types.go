package http

type handler func()

type register struct {
	method string
	path string
	handler handler
}

type Http struct {
	registers []register
}