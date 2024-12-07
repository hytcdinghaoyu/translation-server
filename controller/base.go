package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type BaseController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

var (
	OpSuccess = mvc.Response{Object: map[string]interface{}{"status": 0, "msg": "SUCCESS"}}
)

func (c *BaseController) fireError(err error) mvc.Response {
	return mvc.Response{
		Object: map[string]interface{}{"status": -1, "msg": err.Error()},
	}
}

func (c *BaseController) pageData(data interface{}, total int) mvc.Response {
	ret := make(map[string]interface{})
	ret["data"] = data
	ret["code"] = 0
	ret["count"] = total

	return mvc.Response{
		Object: ret,
	}
}
