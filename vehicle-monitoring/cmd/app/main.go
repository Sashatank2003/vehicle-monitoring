// cmd/app/main.go
package main

import (
	"Sashatank2003/pkg/websocket"
	"log"
	"net/http"
)

func main() {
	// Обработчик WebSocket
	http.HandleFunc("/ws", websocket.HandleWebSocket)

	// Обработчик статики (для HTML, JS, CSS)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	// Запуск сервера
	log.Println("Запуск сервера на порту 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
