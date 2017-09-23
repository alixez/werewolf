package werewolf

import (
	"net/http"
)

type Controller struct {
	context *Context
}

type ControllerInterface interface {
	Init()
	GetContext() *Context
	APISuccess() error
}

func (this *Controller) Init(context *Context) {
	this.context = context
}

func (this *Controller) GetContext() *Context {
	return this.context
}

func (this *Controller) APISuccess(data interface{}) error {
	this.context.apiResponse.Code = 1
	this.context.apiResponse.SubCode = "default.void.success"
	this.context.apiResponse.Message = "接口请求成功"
	if data != nil {

		this.context.apiResponse.Data = data
	} else {
		this.context.apiResponse.Data = map[string]string{}
	}

	return this.context.JSON(http.StatusOK, this.context.apiResponse)
}
