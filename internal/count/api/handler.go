package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	msg "github.com/ValeryBMSTU/web-rk2/internal"
	"github.com/labstack/echo/v4"
)

func (srv *Server) GetTasks(e echo.Context) error {
	var tasks []msg.TaskRepsonse
	var err error
	tasks, err = srv.uc.AllTasks()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, tasks)
}

func (srv *Server) CreateTask(e echo.Context) error {
	req := msg.Task{}
	if err := e.Bind(&req); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	msg, err := srv.uc.CreateTask(req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, msg)
	}
	return e.JSON(http.StatusCreated, msg)
}

func (srv *Server) UpdateTask(e echo.Context) error {
	status := e.QueryParam("status")
	id := e.QueryParam("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return e.JSON(http.StatusBadRequest, err.Error())
		}
		return e.JSON(http.StatusBadRequest, "Произошла ошибка")
	}
	msg, done := srv.uc.UpdateTask(int_id, status)
	if !done {
		return e.String(http.StatusBadRequest, msg)
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) ClearTasks(e echo.Context) error {
	msg, err := srv.uc.ClearTasks()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}
