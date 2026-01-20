package counter_three

import (
	"fmt"
)

func Counter1() {
	var s string
	fmt.Scanln(&s)
	s_byte := []byte(s)
	var counter = 0
	for i := 0; i != len(s_byte); i++ {
		if s_byte[i] == 'a' {
			for j := i; j != len(s_byte); j++ {
				if s_byte[j] == 'b' {
					for k := j; k != len(s_byte); k++ {
						if s_byte[k] == 'c' {
							counter++
						}
					}
				}
			}
		}
	}
	fmt.Println(counter)
}

func Counter2() {
	var s string
	fmt.Scanln(&s)
	s_byte := []byte(s)
	var counter_a, counter_ab, counter_abc = 0, 0, 0
	for _, char := range s_byte {
		switch char {
		case 'a':
			counter_a++
		case 'b':
			counter_ab += counter_a
		case 'c':
			counter_abc += counter_ab
		}
	}
	fmt.Println(counter_abc)
}
