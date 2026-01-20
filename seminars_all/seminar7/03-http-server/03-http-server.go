package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Выводим информацию о запросе
    fmt.Fprintf(w, "Привет! Вы зашли по адресу: %s\n", r.URL.Path)
    fmt.Println("Обработан запрос:", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler) // Обработчик для всех путей
    fmt.Println("Запуск HTTP-сервера на порту 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Ошибка запуска сервера:", err)
    }
}