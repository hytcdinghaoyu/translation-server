package controller

import (
	"encoding/json"
	"fg-admin/config"
	models "fg-admin/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"io"
	"io/ioutil"
	"os"
)

type TaskController struct {
	BaseController
}


func fireError(err error) mvc.Response {
	return mvc.Response{
		Object: map[string]interface{}{"status": -1, "msg": err.Error()},
	}
}

// CreateTask 创建翻译任务
func TaskCreate(ctx iris.Context) mvc.Response {

	file, info, err := ctx.FormFile("file")
	if err != nil {
		return fireError(err)
	}

	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./upload/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fireError(err)
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err == nil {
		config.Reload()
	}

	content, err := ioutil.ReadAll(out)

	data := map[string]interface{}{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{"code": -1, "msg": err.Error()},
		}
	}


	// create new task
	task := new(models.Task)
	task.RawContent = string(content)
	newTask := models.CreateTask(task)


	return mvc.Response{
		Object: map[string]interface{}{"task_id": newTask.ID},
	}

}

func TaskTranslate(ctx iris.Context) {

	//todo: request external LLM API

}

func TaskStatus(ctx iris.Context) {

	//todo: query status from db

}

func TaskDownload(ctx iris.Context) {

	//todo: query translate content from db

}



