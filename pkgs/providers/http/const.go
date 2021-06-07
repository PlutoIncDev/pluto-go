package http

type Method string

const (
	GetMethod Method = "GET"
	PostMethod Method = "POST"
	PatchMethod Method = "PATCH"
	DeleteMethod Method = "DELETE"
	OptionsMethod Method = "OPTIONS"
)