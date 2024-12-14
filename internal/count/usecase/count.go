package usecase

import (
	msg "github.com/ValeryBMSTU/web-rk2/internal"
)

func (u *Usecase) AllTasks() ([]msg.TaskRepsonse, error) {
	var tasks []msg.TaskRepsonse
	var err error
	tasks, err = u.p.SelectAll()
	return tasks, err
}

func (u *Usecase) CreateTask(task msg.Task) (string, error) {
	done, err := u.p.InsertQuery(task)
	if !done {
		return "Произошла ошибка", err
	}
	return "Задача создана", nil
}

func (u *Usecase) UpdateTask(id int, new_status string) (string, bool) {
	status, _ := u.p.SelectStatusByID(id)
	if status == "error" {
		return "Задача не найдена", false
	}
	if status == "new" {
		if new_status != "done" && new_status != "in progress" {
			return "Статус неверный", false
		}
	}
	if status == "in progress" {
		if new_status != "done" {
			return "Статус неверный", false
		}
	}
	done, _ := u.p.UpdateQuery(id, new_status)
	if !done {
		return "Произошла ошибка", false
	}
	return "Задача обновлена", true
}

func (u *Usecase) ClearTasks() (string, error) {
	done, err := u.p.DeleteQuery()
	if !done {
		return "Произошла ошибка", err
	}
	return "Задачи удалены", nil
}
