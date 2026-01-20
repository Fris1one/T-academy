package main

import (
	"fmt"
	"errors"
)

func main() {
	add := func(a, b int64) int64 { //присвоили ПЕРЕМЕННОЙ функцию
		return a+b
	}
	fmt.Println(add(2,3)) //передаем переменную с аргументами
	res := add(2, 3)
	fmt.Println(res)
	result, err := divide(10, 0)
	fmt.Println(result)
	fmt.Println(err) 
}

// функция-объект первого порядка(можно делать тоже самое, что и для других типов данных)
func add(a, b int64) int64 {
	return a + b
}

func swap(a, b int64) (x, y int64) {
	x, y = b, a
	return //можно не писать какие переменные возвращаем, тк это указано в шапке функции(лучше возвращать явно)
}


type MathFunc func(int, int) int //создали свой тип для функции(тоже самое мы делали для переменной в предыдущем семинаре)

func apply(f MathFunc, a, b int) int { //описание функции с функцией внутри
	return f(a,b)
}
func apply1(f func(int, int) int, a, b int) int { //аналогично тому что сверху
	return f(a,b)
}

var crazy = func(x,y int) int {
	return(x+y)*2
}

func crazy1(x,y int) int { //тоже самое что и сверху
	return(x+y)*2
}

func divide(a,b float64) (float64, error) {
	if b == 0{
		return 0, errors.New("Деление на ноль") //проверка, думаю всё понятно
	}
	return a/b, nil
}
