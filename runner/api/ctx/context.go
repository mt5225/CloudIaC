// Copyright 2021 CloudJ Company Limited. All rights reserved.

package ctx

import (
	"cloudiac/runner"
	"cloudiac/utils"
	"cloudiac/utils/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context

	Logger logs.Logger
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context: ctx,
		Logger:  logs.Get().WithField("request", utils.GenPasswd(6, "num")),
	}
}

func (ctx *Context) Result(result interface{}) {
	ctx.JSON(http.StatusOK,
		runner.Response{
			Result: result,
		})
}

func (ctx *Context) Error(err error, code int) {
	ctx.Logger.Errorln(err)
	ctx.JSON(code, runner.Response{
		Error: err.Error(),
	})
}

func HandlerWrapper(handler func(*Context)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		handler(NewContext(ctx))
	}
}
