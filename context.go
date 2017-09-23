package werewolf

import (
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	services    map[string]ServiceInterface
	apiResponse *APIResponse
	dbHelper    map[string]interface{}
	Config      *Env
}

func (this *Context) AddDBHelper(name string, value interface{}) {
	this.dbHelper[name] = value
}

func (this *Context) GetDB(name string) interface{} {
	return this.dbHelper[name]
}

func (this *Context) SetServices(services map[string]ServiceInterface) {

	this.services = services
}

func (this *Context) GetService(name string) ServiceInterface {
	service := this.services[name]
	service.Init(this)
	return service
}
