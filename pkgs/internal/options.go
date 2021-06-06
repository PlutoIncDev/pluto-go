package internal

import "github.com/mcuadros/go-defaults"

type Options struct {
	name string `default:"default"`
	httpPort string `default:"8080"`
}

func NewOptions() *Options {
	opts := new(Options)
	defaults.SetDefaults(opts)
	return opts
}

func (o *Options) SetName(name string) {
	o.name = name
}

func (o *Options) GetName() string {
	return o.name
}

func (o *Options) SetHTTPPort(port string) {
	o.httpPort = port
}









