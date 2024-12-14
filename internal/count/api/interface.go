package api

import msg "github.com/ValeryBMSTU/web-rk2/internal"

type Usecase interface {
	AllTasks() ([]msg.TaskRepsonse, error)
	CreateTask(msg.Task) (string, error)
	UpdateTask(id int, new_status string) (string, bool)
	ClearTasks() (string, error)
}
