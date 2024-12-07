package main

import (
	"fg-admin/config"
	"fg-admin/controller"
	"fg-admin/middleware"
	models "fg-admin/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"log"
)

func main() {
	app := iris.New()

	// 同步数据表
	models.DB.AutoMigrate(
		&models.User{},
	)

	iris.RegisterOnInterrupt(func() {
		_ = models.DB.Close()
	})

	// 注册静态文件目录
	app.HandleDir("/server/upload", "./upload")

	app.Logger().SetLevel("debug")
	// app.Logger().AddOutput()

	// 中间件
	app.Use(recover.New())
	// 敏感操作日志，记入数据库
	app.Use(logger.New())
	app.Use(middleware.NewAuth())

	taskAPI := app.Party("/tasks")
	{
		// POST: http://localhost:8080/tasks
		taskAPI.Post("/", controller.TaskCreate)
		// POST: http://localhost:8080/tasks/{task_id}/translate
		taskAPI.Post("/{task_id}/translate", controller.TaskTranslate)
		// GET: http://localhost:8080/tasks/{task_id}
		taskAPI.Get("/{task_id}", controller.TaskStatus)
		// GET: http://localhost:8080/tasks/{task_id}/download
		taskAPI.Get("/{task_id}/download", controller.TaskDownload)

	}

	mvc.Configure(app.Party("/"), basicMVC)
	err := app.Run(iris.Addr(config.Conf.Get("app.addr").(string)), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		log.Fatal(err)
	}
}

func basicMVC(app *mvc.Application) {
	app.Register(
		sessions.New(sessions.Config{}).Start,
	)

	app.Party("/auth").Handle(new(controller.UserController))

}
