package http

import (
	"github.com/gin-gonic/gin"
)

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		HTTP: ctx,
	}
}
