package websocket

import (
    "manageme/internal/models"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
)

var (
    clients   = make(map[*websocket.Conn]bool)
    clientsMu sync.Mutex

    upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
        return
    }

    clientsMu.Lock()
    clients[conn] = true
    clientsMu.Unlock()

    defer func() {
        clientsMu.Lock()
        delete(clients, conn)
        clientsMu.Unlock()
        conn.Close()
    }()

    for {
        _, _, err := conn.ReadMessage()
        if err != nil {
            break
        }
    }
}

func NotifyAllClients(action string, task models.Task) {
    clientsMu.Lock()
    defer clientsMu.Unlock()

    for client := range clients {
        err := client.WriteJSON(map[string]interface{}{
            "action": action,
            "task":   task,
        })
        if err != nil {
            client.Close()
            delete(clients, client)
        }
    }
}
