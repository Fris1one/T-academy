package strange_row

import (
	"fmt"
)

func Strange() {
	var length int
	fmt.Scanln(&length)
	var s string
	fmt.Scanln(&s)
	var slice = []byte(s)
	var result_slice = make([]int, 0, length+1)
	result_slice = append(result_slice, 0)
	var ptr = 0
	for i:= 0; i != length; i++ {
		if i == 0 {
		}
		fmt.Println(result_slice)
		if slice[i] == 'R' {
			result_slice = append(result_slice[:ptr+1], append([]int{i+1}, result_slice[ptr+1:]...)...)
			ptr++
		} else {
			result_slice = append(result_slice[:ptr], append([]int{i+1}, result_slice[ptr:]...)...)
		}
	} 
	for i:= range result_slice {
        fmt.Print(i, " ")
    }
} 