package controller

import (
	"errors"
	"fg-admin/model"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	BaseController
}


// PostUsers 注册新用户
func (c *UserController) PostUsers() mvc.Result {
	username := c.Ctx.PostValueTrim("username")
	pwd := c.Ctx.PostValueTrim("password")

	user := models.GetUserByUserName(username)
	if user.ID == 0 {
		// create new user
		user.Username = username
		user.Password = pwd
		models.CreateUser(user)
	} else {
		return mvc.Response{Code: -1, Err: errors.New("username exists!")}
	}

	return  OpSuccess
}

// PostLogin 登录请求
func (c *UserController) PostLogin() mvc.Result {
	uname := c.Ctx.PostValueTrim("username")
	pwd := c.Ctx.PostValueTrim("password")

	token, ok, msg := models.CheckPwd(uname, pwd)
	if ok {
		return mvc.Response{Code: -1, Err: errors.New(msg)}
	}

	ret := map[string]interface{}{"token": token}
	return mvc.Response{
		Object: ret,
	}
}
