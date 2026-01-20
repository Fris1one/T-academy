package main

import (
	"fmt"
	// "net/http"
)

func main() {
	increment(decrement(final()))(func() {})

}

type middleware func(func())

func increment(m middleware) middleware {
	return func(f func()) {
		fmt.Println("Before increment") 

		m(f)

		fmt.Println("After increment")
	}
}
func decrement(m middleware) middleware {
	return func(f func()) {
		fmt.Println("Before decrement") 

		m(f)

		fmt.Println("After decrement")
	}
}

func final() middleware {
	return func(f func()) {
		fmt.Println("Final")
	}
}