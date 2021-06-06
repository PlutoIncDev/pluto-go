package pluto

import (
	"fmt"
	"pluto/pkgs/internal"
	"pluto/pkgs/logging"
	"sync"
)

type Client struct {
	Options *internal.Options
	Events *internal.Events
}

func NewClient() *Client {
	return &Client{
		Options: internal.NewOptions(),
		Events: internal.NewEvents(),
	}
}

func (c *Client) String() string {
	return fmt.Sprintf("<PlutoClient name='%s'/>", c.Options.GetName())
}

func (c *Client) Start() error {
	var wg sync.WaitGroup
	defer wg.Wait()

	logging.Info(fmt.Sprintf("Starting '%s' Service 🚀", c.Options.GetName()))

	// Start the HTTP Server (if necessary)
	wg.Add(1)
	go startHttpServer(&wg, c)

	// Start the AMQP Connection
	wg.Add(1)
	go startAMQPConnection(&wg, c)

	return nil
}
