// web/script.js
let socket;

document.addEventListener("DOMContentLoaded", function() {
    // Открытие WebSocket-соединения
    socket = new WebSocket("ws://localhost:8080/ws");

    // Событие при успешном подключении
    socket.onopen = function(event) {
        console.log("WebSocket соединение установлено");
    };

    // Событие при получении сообщения
    socket.onmessage = function(event) {
        const message = event.data;
        const messagesDiv = document.getElementById("messages");
        messagesDiv.innerHTML += `<p>Получено сообщение: ${message}</p>`;
    };

    // Событие при закрытии соединения
    socket.onclose = function(event) {
        console.log("WebSocket соединение закрыто");
    };

    // Событие при ошибке
    socket.onerror = function(error) {
        console.error("Ошибка WebSocket:", error);
    };

    // Отправка сообщения при нажатии на кнопку
    document.getElementById("sendMessageBtn").addEventListener("click", function() {
        const message = "Привет, сервер!";
        socket.send(message);
        console.log("Отправлено сообщение:", message);
    });
});
