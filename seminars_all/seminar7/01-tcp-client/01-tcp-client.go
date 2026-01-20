package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Подключение к серверу по адресу localhost:8080
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Ошибка при подключении:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Можно ввести сообщение с клавиатуры
    fmt.Print("Введите сообщение для отправки: ")
    reader := bufio.NewReader(os.Stdin)
    message, _ := reader.ReadString('\n')

    // Отправляем сообщение
    fmt.Fprintf(conn, "%s", message)

    // Читаем ответ от сервера
    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Ошибка при чтении ответа:", err)
        return
    }
    fmt.Println("Ответ от сервера:", response)
}