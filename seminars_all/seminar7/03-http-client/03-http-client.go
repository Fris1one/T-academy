package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("http://localhost:8080/")
    if err != nil {
        fmt.Println("Ошибка при запросе:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Ошибка при чтении ответа:", err)
        return
    }

    fmt.Println("Ответ сервера:")
    fmt.Println(string(body))
}