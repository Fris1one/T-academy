package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // Создаем UDP-слушатель на порту 8080
    addr := net.UDPAddr{
        Port: 8080,
        IP:   net.ParseIP("0.0.0.0"),
    }

    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Println("Ошибка при запуске UDP-сервера:", err)
        os.Exit(1)
    }
    defer conn.Close()

    fmt.Println("UDP-сервер запущен и слушает порт 8080")

    buffer := make([]byte, 1024)

    for {
        // Читаем входящее сообщение
        n, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Ошибка при чтении данных:", err)
            continue
        }

        receivedMessage := string(buffer[:n])
        fmt.Printf("Получено сообщение от %s: %s\n", remoteAddr.String(), receivedMessage)

        // Отправляем ответ обратно клиенту
        response := "Сообщение получено: " + receivedMessage
        _, err = conn.WriteToUDP([]byte(response), remoteAddr)
        if err != nil {
            fmt.Println("Ошибка при отправке ответа:", err)
        }
    }
}