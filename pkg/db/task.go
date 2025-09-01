package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

var ErrTaskNotFound = errors.New("task not found")

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	// определите запрос
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	row, err := DB.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?",
		limit,
	)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var tasks = make([]*Task, 0, limit)
	for row.Next() {
		var task Task
		if err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	row := DB.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?",
		idInt,
	)

	var task Task
	err = row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for updating task")
	}
	return nil
}

func DeleteTask(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	res, err := DB.Exec("DELETE FROM scheduler WHERE id = ?", idInt)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrTaskNotFound
	}
	return nil
}
