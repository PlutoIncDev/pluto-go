package http

import (
	"fmt"
	"pluto/pkgs/errors"
	"pluto/pkgs/logging"
)


func (h Http) HTTP(method string, path string, handler handler) {

	// check method is valid
	if !h.isValidHttpMethod(method) {
		logging.Panic(errors.NewErr(fmt.Sprintf("%s is not a supported HTTP method", method)))
	}

	r := register{
		method: method,
		path:    path,
		handler: handler,
	}

	// check this path hasn't already been registered for this method
	if !h.isNewRegister(r) {
		logging.Panic(errors.NewErr(fmt.Sprintf("path=%s is registered more than once", r.path)))
	}

	h.registers = append(h.registers, r)
	logging.Info("HTTP registered method", r.method, "for path", r.path, len(h.registers))
}


