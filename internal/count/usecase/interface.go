package usecase

import msg "github.com/ValeryBMSTU/web-rk2/internal"

type Provider interface {
	SelectStatusByID(id int) (string, error)
	SelectAll() ([]msg.TaskRepsonse, error)
	InsertQuery(task msg.Task) (bool, error)
	UpdateQuery(id int, status string) (bool, error)
	DeleteQuery() (bool, error)
}
