package client

import (
	"context"
	"fmt"
	"log"
	"pluto/pkgs/providers/base"
	"sync"
	"time"
)

type Client struct {
	// config
	name                    string
	maxShutdownDuration     time.Duration // max duration for shutting down all the providers when the server closes before force quitting
	isFinishedSleepDuration time.Duration // how often to poll the services when shutting down to check if they're ready to stop

	// internal variables
	providers     []base.Provider
	cancelCtxFunc context.CancelFunc
}

func New(name string) *Client {
	return &Client{
		name:                    name,
		maxShutdownDuration:     time.Second * 30,
		isFinishedSleepDuration: time.Second,
	}
}

func (c *Client) Start(blocking bool) {
	log.Println(fmt.Sprintf("Starting the %s service", c.name))

	ctx, cancel := context.WithCancel(context.Background())
	c.cancelCtxFunc = cancel // assign for later use
	var wg sync.WaitGroup

	for _, p := range c.providers {
		wg.Add(1)

		provider := p
		go func() {
			go provider.Run(ctx)

			for {
				select {
				case <-ctx.Done():
					provider.Shutdown()
					return
				}
			}
		}()
	}

	if blocking {
		wg.Wait()
	}
}

func (c *Client) Stop() {
	log.Println("Attempting to stop service")
	c.cancelCtxFunc()

	log.Println("Issued STOP signal to providers")

	startShutdownTime := time.Now()

	// check if all providers are finished
	for {
		if time.Now().Sub(startShutdownTime) > c.maxShutdownDuration {
			log.Println("Some providers took longer to shutdown than the maxShutdownDuration")
			break
		}

		// check if providers are ready to shut down every second
		secsRemaining := int((c.maxShutdownDuration - time.Now().Sub(startShutdownTime)).Seconds())

		log.Println(fmt.Sprintf("Waiting for providers to shutdown (%ds remaining before force quitting)", secsRemaining))
		time.Sleep(c.isFinishedSleepDuration)

		providerStillRunning := false
		for _, p := range c.providers {
			if !p.IsFinished() {
				providerStillRunning = true
			}
		}

		if !providerStillRunning {
			// if we reached here, all the providers are finished
			break
		}
	}

	log.Println("Shutdown service complete")
}

func (c *Client) RegisterProvider(p base.Provider) {
	// todo: stop providers being registered twice reflect type/ name maybe?
	c.providers = append(c.providers, p)
	p.Setup()
}
