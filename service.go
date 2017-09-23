package werewolf

type Service struct {
	model   string
	context *Context
}

func (this *Service) Init(ctx *Context) {
	this.context = ctx
}

func (this *Service) GetModel() string {
	return this.model
}

type ServiceInterface interface {
	Init(context *Context)
	GetModel() string
}
