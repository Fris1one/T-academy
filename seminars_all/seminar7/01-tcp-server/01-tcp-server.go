package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Слушаем порт 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Ошибка при запуске сервера:", err)
        os.Exit(1)
    }
    defer listener.Close()
    fmt.Println("Сервер запущен и слушает порт 8080")

    for {
        // Принимаем входящее соединение
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Ошибка при принятии соединения:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Клиент подключился:", conn.RemoteAddr())

    reader := bufio.NewReader(conn)
    // Читаем сообщение от клиента
    message, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Ошибка при чтении сообщения:", err)
        return
    }
    fmt.Printf("Получено сообщение от клиента: %s", message)

    // Отправляем ответ
    response := "Сообщение получено: " + message
    conn.Write([]byte(response))
}