package pluto

import (
	"fmt"
	"log"
	"pluto/pkgs/providers"
	"pluto/pkgs/providers/base"
	"sync"
)

type Client struct {
	name string
	providers []base.Provider
}

func NewClient(name string) *Client {
	// TODO: replace name with config (and validate config)
	if len(name) == 0 {
		panic("Name of service must be at least 1 character long")
	}

	return &Client{name: name}
}

func (c *Client) RegisterProvider(p base.Provider) {
	c.providers = append(c.providers, p)

	p.Setup()
}

func (c *Client) GetProviders() []base.Provider {
	return c.providers
}

func (c *Client) Start() {
	log.Println(fmt.Sprintf("Starting the %s service", c.name))

	var wg sync.WaitGroup

	for _, p := range c.providers {
		wg.Add(1)
		go providers.RunProvider(&wg, p)
	}

	wg.Wait()

	for _, p := range c.providers {
		p.Shutdown()
	}
}
