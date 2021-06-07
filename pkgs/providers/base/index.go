package base

import "context"

type Provider interface {
	Setup()
	Run(ctx context.Context)
	Shutdown()
	IsFinished() bool // should be true when the provider is ready to be finished
}
