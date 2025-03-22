// pkg/websocket/websocket.go
package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Инициализация upgrader для WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешить все подключения (для разработки)
	},
}

// Обработчик WebSocket
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем WebSocket-соединение
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка при установке WebSocket-соединения:", err)
		return
	}
	defer conn.Close()

	// Чтение сообщений от клиента
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения:", err)
			break
		}
		// Логирование полученного сообщения
		log.Printf("Получено сообщение: %s", message)

		// Отправка сообщения обратно клиенту
		err = conn.WriteMessage(websocket.TextMessage, []byte("Ответ от сервера: "+string(message)))
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
			break
		}
	}
}
