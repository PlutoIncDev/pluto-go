package pluto

import (
	"fmt"
	"log"
	"pluto/pkgs/providers/base"
	"sync"
	"context"
)

type Client struct {
	name      string
	providers []base.Provider
	ctx context.Context
	ctxCancel context.CancelFunc
}

func NewClient(name string) *Client {
	// TODO: replace name with config (and validate config)
	if len(name) == 0 {
		panic("Name of service must be at least 1 character long")
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Client{name: name, ctx: ctx, ctxCancel: cancel}
}

func (c *Client) RegisterProvider(p base.Provider) {

	c.providers = append(c.providers, p)

	p.Setup()
}

func (c *Client) GetProviders() []base.Provider {
	return c.providers
}

func (c *Client) Start(async bool) {
	log.Println(fmt.Sprintf("Starting the %s service [async=%t]", c.name, async))

	var wg sync.WaitGroup

	for _, p := range c.providers {
		wg.Add(1)

		p := p // loop variable capture issue
		go func() {
			defer wg.Done()

			select {
			case <-c.ctx.Done():
				log.Println("DONE!!!")
				return
			default:

			}

			p.Run()
		}()
	}

	if !async {
		wg.Wait()
		// todo: handle sigterm?
	}
}

func (c *Client) Stop() {
	log.Println(fmt.Sprintf("Stopping the %s service", c.name))
	c.ctxCancel()

	//for _, cp := range c.providers {
	//	log.Println(cp.channel)
	//	cp.channel <- 1
	//	close(cp.channel)// closes the service
	//}
}