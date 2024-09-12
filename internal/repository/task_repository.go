package repository

import (
    "database/sql"
    "manageme/internal/models"
)

func CreateTask(db *sql.DB, task *models.Task) (string, error) {
    var id string
    err := db.QueryRow(
        "INSERT INTO tasks (title, description, status, due_date) VALUES ($1, $2, $3, $4) RETURNING id",
        task.Title, task.Description, task.Status, task.DueDate,
    ).Scan(&id)
    return id, err
}

func UpdateTask(db *sql.DB, id string, task *models.Task) (models.Task, error) {
    _, err := db.Exec(
        "UPDATE tasks SET title = $1, description = $2, status = $3, due_date = $4 WHERE id = $5",
        task.Title, task.Description, task.Status, task.DueDate, id,
    )
    return models.Task{ID: id, Title: task.Title, Description: task.Description, Status: task.Status, DueDate: task.DueDate}, err
}

func DeleteTask(db *sql.DB, id string) error {
    _, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
    return err
}

func GetAllTasks(db *sql.DB) ([]models.Task, error) {
    rows, err := db.Query("SELECT id, title, description, status, due_date FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}
