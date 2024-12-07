package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	RawContent string
	TranslateContent string
	Status   uint
}

func FindTaskByID(id int64) Task {
	u := Task{}
	if err := DB.Where("id = ?", id).First(&u).Error; err != nil {
		fmt.Printf("DB Error:%s", err)
	}
	return u
}

func CreateTask(t *Task) (task *Task) {

	task = new(Task)
	task.RawContent = t.RawContent
	task.Status = 0

	if err := DB.Create(task).Error; err != nil {
		fmt.Printf("CreateTaskErr:%s", err)
	}

	return
}
