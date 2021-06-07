package main

import (
	"github.com/gin-gonic/gin"
	"pluto/pkgs/client"
	"pluto/pkgs/providers/http"
)

func main() {
	c := client.New("test-client")

	httpProvider := http.NewProvider("8080")
	httpProvider.RegisterEndpoint("GET", "/health-check", func(ctx *http.Context) {
		ctx.HTTP.JSON(200, gin.H{"test": "calum"})
	})

	c.RegisterProvider(httpProvider)

	c.Start(true)

	//time.Sleep(time.Second * 10)
	//c.Stop()
}
