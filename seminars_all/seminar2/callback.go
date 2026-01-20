package main

import (
	"fmt"
	"os"
)

func main() {

}

type FileResource func(cont FileCallback) error

type FileCallback func(fd *os.File) error

//аункция для бизнес-логики
func WorkWithFile(path string, flags int, perm os.FileMode) FileResource {
	//каррированный инициализатор ресурса
	return func(cb FileCallback) error {
		//системные вызовы
		file, err := os.OpenFile(path, flags, perm)

		if err != nil {
			return err
		}

		fmt.Println("File opened")

		defer func() {
			_ = file.Close()

			fmt.Println("File closed")
		}()

		//сама бизнес-логика(пользовательская)
		err = cb(file)

		return err
	}

}