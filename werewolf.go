package werewolf

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	Application struct {
		Echo        *echo.Echo
		Controllers map[string]interface{}
		Services    map[string]interface{}
		Router      *Router
	}

	BootCallBackFunc func(application *Application) error
)

func (this *Application) getType(typeOf interface{}) reflect.Type {
	return reflect.Indirect(reflect.ValueOf(typeOf)).Type()
}

func (this *Application) AddController(controller interface{}) {
	fmt.Println(controller)
	fmt.Println(this.getType(controller).Name())
	this.Controllers[this.getType(controller).Name()] = controller
}

func (this *Application) AddService(service interface{}) {
	this.Services[this.getType(service).Name()] = service
}

func (this *Application) initRouter() {
	this.Router.ControllersIndex = this.Controllers
}

func (this *Application) Boot(callback BootCallBackFunc) {
	callback(this)
	this.initRouter()
}

func (this *Application) Start(address string) {
	this.Echo.Start(address)
}

func BetterAppContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{
			c,
			nil,
			&APIResponse{
				Code:    0,
				Message: "空",
				SubCode: "default.void.default",
				Data:    map[string]interface{}{},
			},
			map[string]interface{}{},
		}

		return next(cc)
	}
}

func AddGormToContext(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c := ctx.(*Context)
			c.AddDBHelper("gorm", db)
			return next(c)
		}
	}
}

func CreateApplication(env *Env) (application *Application) {
	e := echo.New()
	e.Use(BetterAppContext)

	dbConfig := env.GetConfig("database").(map[interface{}]interface{})
	dbConfigStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", dbConfig["user"], dbConfig["password"], dbConfig["host"].(string)+":"+dbConfig["port"].(string), dbConfig["db"], dbConfig["charset"])
	db, err := gorm.Open("mysql", dbConfigStr)

	if err != nil {
		application.Echo.Logger.Fatal(err)
		fmt.Println(err)
	} else {
		application.Echo.Logger.Debug("数据库连接成功!")
		fmt.Println("(: 数据库连接成功 ...")
	}
	e.Use(AddGormToContext(db))

	router := &Router{
		Echo: e,
	}
	application = &Application{
		Echo:        e,
		Router:      router,
		Controllers: make(map[string]interface{}),
		Services:    make(map[string]interface{}),
	}
	return
}
