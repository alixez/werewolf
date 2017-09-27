package werewolf

import (
	"reflect"
	"strings"

	"github.com/labstack/echo"
)

func HandleFunc(ctx echo.Context, action string, method string, controllerIndex map[string]interface{}) error {
	c := ctx.(*Context)
	if controllerIndex[action] == nil {
		return echo.ErrNotFound
		// return errors.New("Not Found the action")
	}
	controllerType := reflect.Indirect(reflect.ValueOf(controllerIndex[action])).Type()
	vc := reflect.New(controllerType)
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(c)
	vc.MethodByName("Init").Call(in)
	res := vc.MethodByName(method).Call(make([]reflect.Value, 0))
	if len(res) == 0 {
		return nil
	}
	if res[0].Interface() == nil {
		return nil
	}

	return res[0].Interface().(error)

}

func ParseActionStr(str string) ([]string, error) {
	arr := strings.Split(str, "@")

	if len(arr) != 2 {
		return nil, echo.ErrNotFound
	}

	return arr, nil
}

type (
	Router struct {
		Echo             *echo.Echo
		ControllersIndex map[string]interface{}
	}

	Group struct {
		echoGroup *echo.Group
		router    *Router
	}
)

func (g *Group) Group(perfix string, m ...echo.MiddlewareFunc) *Group {
	echoGroup := g.echoGroup.Group(perfix, m...)

	group := &Group{
		echoGroup,
		g.router,
	}

	return group
}

func (g *Group) Get(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	g.echoGroup.GET(url, func(c echo.Context) error {
		if err != nil {
			return err
		}

		return HandleFunc(c, actionArr[0], actionArr[1], g.router.ControllersIndex)
	}, m...)
}

func (g *Group) Post(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	g.echoGroup.POST(url, func(c echo.Context) error {
		if err != nil {
			return err
		}

		return HandleFunc(c, actionArr[0], actionArr[1], g.router.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Group(perfix string, m ...echo.MiddlewareFunc) *Group {
	echoGroup := appRoute.Echo.Group(perfix, m...)

	group := &Group{
		echoGroup,
		appRoute,
	}

	return group
}

func (appRoute *Router) Get(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.GET(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Post(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.POST(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Delete(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.DELETE(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Put(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.PUT(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Patch(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.PATCH(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Any(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.Any(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Trace(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.TRACE(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}

func (appRoute *Router) Options(url string, actionStr string, m ...echo.MiddlewareFunc) {
	actionArr, err := ParseActionStr(actionStr)
	appRoute.Echo.OPTIONS(url, func(c echo.Context) error {
		if err != nil {
			return err
		}
		return HandleFunc(c, actionArr[0], actionArr[1], appRoute.ControllersIndex)
	}, m...)
}
