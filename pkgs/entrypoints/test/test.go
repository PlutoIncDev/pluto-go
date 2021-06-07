package test

import "pluto/pkgs/logging"

type TestEntrypoint struct {

}

func NewTestEntrypoint() TestEntrypoint {
	return TestEntrypoint{}
}

func (te TestEntrypoint) Setup() {
	logging.Info("Setup Test Called")
}

func (te TestEntrypoint) Shutdown() {
	logging.Info("Shutdown Test Called")
}

func (te TestEntrypoint) Run() {
	logging.Info("Run Test")
}