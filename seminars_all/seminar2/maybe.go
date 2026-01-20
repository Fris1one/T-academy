package main
import (
	
)
//это дженерики
func main() {

}

type Maybe[T any] struct {
	value *T
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value: &value }
}