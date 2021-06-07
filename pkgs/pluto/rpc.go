package pluto

import (
	"pluto/pkgs/logging"
	"sync"
)

func startRpcConnection(wg *sync.WaitGroup, c *Client) {
	defer wg.Done()

	handlers := c.Events.Rpc.GetHandlers()

	if len(handlers) <= 0 {
		logging.Info("No RPC handlers registered")
		return
	}

	logging.Info("Starting the RPC Server with", len(handlers), "handlers")
}
