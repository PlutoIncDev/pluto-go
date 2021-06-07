package http

import (
	"pluto/pkgs/logging"
)


func New() Http {
	return Http{}
}

func (h *Http) Setup() {
	logging.Info("Setting up HTTP Entrypoint")
}

func (h *Http) Shutdown() {
	logging.Info("Shutting down HTTP Entrypoint")
}

func (h *Http) Run() {
	logging.Info("Running HTTP Entrypoint")
	//e := gin.New()
}
