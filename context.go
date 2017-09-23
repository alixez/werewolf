package werewolf

import (
	"errors"
	"reflect"

	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	services    map[string]reflect.Value
	apiResponse *APIResponse
	DBHelper    map[string]interface{}
}

func (this *Context) AddDBHelper(name string, value interface{}) {
	this.DBHelper[name] = value
}

func (this *Context) GetDB(name string) interface{} {
	return this.DBHelper[name]
}

func (this *Context) GetService(name string) (error, interface{}) {
	service := this.services[name]
	if service.Interface() == nil {
		return errors.New("Not exists the service"), nil
	}

	return nil, service.Interface()
}
