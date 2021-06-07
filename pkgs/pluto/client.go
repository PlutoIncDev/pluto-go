package pluto

import (
	"fmt"
	"pluto/pkgs/entrypoints"
	"pluto/pkgs/errors"
	"pluto/pkgs/internal"
	"pluto/pkgs/logging"
	"reflect"
	"sync"
)

type Client struct {
	Options *internal.Options
	Events *internal.Events
	Entrypoints []entrypoints.Entrypoint
}

func NewClient(name string) (*Client, error) {
	opts, err := internal.NewOptions(name)
	if err != nil {
		return nil, err
	}

	return &Client{
		Options: opts,
		// todo; error handling on new events?
		Events: internal.NewEvents(),
	}, nil
}

func (c *Client) String() string {
	return fmt.Sprintf("<PlutoClient name='%s'/>", c.Options.GetName())
}

func (c *Client) RegisterEntrypoint(e entrypoints.Entrypoint) {
	c.Entrypoints = append(c.Entrypoints, e)

	e.Setup()
}

func (c *Client) Start() {
	var wg sync.WaitGroup

	logging.Info(fmt.Sprintf("Starting %s Service 🚀", c.Options.GetName()))

	if len(c.Entrypoints) == 0 {
		logging.Panic(errors.NewErr("Unable to start service as no entry points defined"))
	}

	for _, e := range c.Entrypoints {
		logging.Info(fmt.Sprintf("Running Entrypoint %s", reflect.TypeOf(e)))

		wg.Add(1)
		go entrypoints.RunEntrypoint(&wg, e)
	}

	wg.Wait()

	for _, e := range c.Entrypoints {
		e.Shutdown()
	}
}
