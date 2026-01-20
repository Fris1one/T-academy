package main

import (
	"fmt"
	"iter"
	"time"
)

func main() {
	// closures()
	fibonacci()
}
// Замыкание - это функция, которая запоминает значения из своей внешней области видимости,
// даже если эта область уже недоступна. Она создается, когда функция объявляется, и продолжает
// запоминать значения переменных даже после того, как вызывающая функция завершит свою работу.

// Замыкания - это инструмент, который позволяет сохранять значения и состояние между вызовами функций,
// создавать функции на лету и возвращать их из других функций.

//чел сказал, что не использует их, это просто для инфы

func cleanSeq() func() int {
	i := 0
	return func() int { //возвращаемая функцией функция(можно вызывать несколько раз)
		i++
		return i
	}
}

var j = 0

func dirtySeq() func() int { //в грязном замыкании используем глобальную переменную
	return func() int {
		j++
		return j
	}
}

func closures() {
	clean := cleanSeq()

	fmt.Println(clean())
	fmt.Println(clean())
	fmt.Println(clean())
	newInt := cleanSeq()
	for i := 0; i < 10; i++ {
		fmt.Println(newInt())
	}
	dirt := dirtySeq()
	fmt.Println(dirt())
	fmt.Println(dirt())
	fmt.Println(dirt())
	j = -1 //поменяли переменную, от которой зависит func dirtySeq
	fmt.Println(dirt())

}




func fib() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}
func toSlice(n int, f func() int) []int {
	s:= make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, f())
	}
	return s
}
func fibonacci() {
	f := fib()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(toSlice(50, fib()))
}


func intSeqIt() iter.Seq[int] { //чел сказал не вникать в это-_-(это что-то с дженериками)
	return func(yield func(int) bool) {
		for i := 0; i < 100; i++ {
			time.Sleep(100*time.Millisecond)
			if !yield(i) {
				return
			}
		}
	}
}