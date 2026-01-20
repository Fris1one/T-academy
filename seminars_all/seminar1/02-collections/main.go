package main

import "fmt"

// Создайте слайс и заполните его числами от 1 до 100. 
// Оставьте в слайсе первые и последние 10 элементов и разверните слайс в обратном порядке.
// Подумайте, можно ли подобным образом развернуть строку?
func main() {
	var dim = 100
	slice1 := make([]int, 0, dim)
	for i := 0; i < dim; i++ {
		slice1 = append(slice1, i+1)
	}
	slice1 = append(slice1[:10], slice1[dim-10:]...)
	fmt.Println(slice1)
	dim = len(slice1)
	for i := range slice1[:dim/2] {
		slice1[i], slice1[dim-i-1] = slice1[dim-i-1], slice1[i]
	}
	fmt.Println(slice1)


	var a map[string]int = map[string]int{"one":1, "two":2}
	fmt.Println(a)
}



