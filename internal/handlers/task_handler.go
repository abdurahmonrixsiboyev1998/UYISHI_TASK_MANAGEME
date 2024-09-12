package handler

import (
	"database/sql"
	"encoding/json"
	"manageme/internal/models"
	"manageme/internal/repository"
	"manageme/internal/websocket"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.ID, _ = repository.CreateTask(db, &task)
		websocket.NotifyAllClients("created", task)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var updatedTask models.Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, err := repository.UpdateTask(db, id, &updatedTask)
		if err != nil {
			http.Error(w, "No task found", http.StatusNotFound)
			return
		}

		websocket.NotifyAllClients("updated", task)
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTaskHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := repository.DeleteTask(db, id); err != nil {
			http.Error(w, "No task found", http.StatusNotFound)
			return
		}

		websocket.NotifyAllClients("deleted", models.Task{ID: id})
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetTasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := repository.GetAllTasks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tasks)
	}
}
