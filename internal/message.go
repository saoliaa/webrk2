package message

import "time"

type Task struct {
	Author_name   string `json:"author_name"`
	Assignee_name string `json:"assignee_name"`
}

type TaskRepsonse struct {
	ID            int        `json:"id"`
	Author_name   string     `json:"author_name"`
	Assignee_name string     `json:"assignee_name"`
	Status        string     `json:"status"`
	Created_date  time.Time  `json:"created_date"`
	Resolved_date *time.Time `json:"resolved_date"`
}
