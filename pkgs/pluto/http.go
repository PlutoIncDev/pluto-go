package pluto

import (
	"pluto/pkgs/internal"
	"pluto/pkgs/logging"
	"sync"
)

func startHttpServer(wg *sync.WaitGroup, c *Client) {
	defer wg.Done()

	handlers := c.Events.Http.GetHandlers()

	if len(handlers) <= 0 {
		logging.Info("No HTTP handlers registered")
		return
	}

	logging.Info("Starting the HTTP Server with", len(handlers), "handlers")

	for _, handler := range c.Events.Http.GetHandlers() {
		logging.Debug("Registering", handler.Method, "HTTP Method for path", handler.Path)

		switch handler.Method {
		case internal.GetHttpMethod:
			break
		case internal.PostHttpMethod:
			break
		case internal.PatchHttpMethod:
			break
		case internal.DeleteHttpMethod:
			break
		default:
			panic("UNABLE TO REGISTER THIS ONE!")
		}
	}
}
