package internal

import (
	"github.com/mcuadros/go-defaults"
	"pluto/pkgs/errors"
)

type Options struct {
	name string
	httpPort string `default:"8080"`
}

func NewOptions(name string) (*Options, error) {
	// check name is valid
	if len(name) < 3 {
		return nil, errors.NewErr("Name must be at least 3 characters long.")
	}

	opts := &Options{name: name}
	defaults.SetDefaults(opts)
	return opts, nil
}

func (o *Options) GetName() string {
	return o.name
}








