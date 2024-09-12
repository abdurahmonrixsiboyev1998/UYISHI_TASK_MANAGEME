package main

import (
	"database/sql"
	"log"
	"net/http"

	"manageme/config"
	handler "manageme/internal/handlers"

	websocket"manageme/internal/websocket"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	cfg := config.LoadConfig()

	var err error
	db, err = sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ma'lumotlar bazasiga ulanishda xato: %v", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.CreateTaskHandler(db)).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskHandler(db)).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler(db)).Methods("DELETE")
	router.HandleFunc("/api/tasks", handler.GetTasksHandler(db)).Methods("GET")

	router.HandleFunc("/ws", websocket.HandleWebSocket)

	log.Println("Server ishlamoqda: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
