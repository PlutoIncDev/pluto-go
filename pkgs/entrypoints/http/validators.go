package http

import (
	"pluto/pkgs/internal/utils"
	"pluto/pkgs/logging"
)

func (h Http) isValidHttpMethod(method string) bool {
	return utils.StringInSlice(method, validHttpMethods)
}

func (h Http) isNewRegister(r register) bool {
	logging.Info(h.registers, r)
	for _, existingR := range h.registers {

		if r.path == existingR.path && r.method == existingR.method {
			return false
		}
	}

	return true
}