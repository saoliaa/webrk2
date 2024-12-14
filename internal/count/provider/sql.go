package provider

import (
	msg "github.com/ValeryBMSTU/web-rk2/internal"
)

func (p *Provider) SelectStatusByID(id int) (string, error) {
	var status string
	err := p.conn.QueryRow("SELECT status FROM tasks WHERE id = $1", id).Scan(&status)
	if err != nil {
		return "error", err // обработка ошибки
	}

	return status, nil
}

func (p *Provider) SelectAll() ([]msg.TaskRepsonse, error) {
	var tasks []msg.TaskRepsonse

	rows, err := p.conn.Query("SELECT * FROM tasks")
	if err != nil {
		return tasks, err // обработка ошибки
	}
	defer rows.Close() // закроем rows после окончания работы

	for rows.Next() {
		var task msg.TaskRepsonse
		// Здесь нужно указать все поля структуры TaskResponse
		err := rows.Scan(&task.ID, &task.Author_name, &task.Assignee_name, &task.Status, &task.Created_date, &task.Resolved_date)
		if err != nil {
			return tasks, err // обработка ошибки
		}
		tasks = append(tasks, task) // добавление структуры в массив
	}

	return tasks, nil
}

func (p *Provider) InsertQuery(task msg.Task) (bool, error) {
	_, err := p.conn.Exec("INSERT INTO tasks (author_name, assignee_name, status, created_date) VALUES ($1, $2, 'new', CURRENT_DATE)",
		task.Author_name, task.Assignee_name)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Provider) UpdateQuery(id int, status string) (bool, error) {
	_, err := p.conn.Exec("UPDATE tasks SET status = $1 WHERE id = $2;", status, id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Provider) DeleteQuery() (bool, error) {
	_, err := p.conn.Exec("DELETE FROM tasks")

	if err != nil {
		return false, err
	}

	return true, nil
}
