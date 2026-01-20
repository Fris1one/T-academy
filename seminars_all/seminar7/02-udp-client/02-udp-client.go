package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    serverAddr := net.UDPAddr{
        Port: 8080,
        IP:   net.ParseIP("127.0.0.1"),
    }

    conn, err := net.DialUDP("udp", nil, &serverAddr)
    if err != nil {
        fmt.Println("Ошибка при подключении к серверу:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Ввод сообщения с клавиатуры
    fmt.Print("Введите сообщение для отправки: ")
    reader := bufio.NewReader(os.Stdin)
    message, _ := reader.ReadString('\n')

    // Отправка сообщения серверу
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Ошибка при отправке:", err)
        return
    }

    // Чтение ответа
    buffer := make([]byte, 1024)
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        fmt.Println("Ошибка при получении ответа:", err)
        return
    }

    response := string(buffer[:n])
    fmt.Println("Ответ от сервера:", response)
}