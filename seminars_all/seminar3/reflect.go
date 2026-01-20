package main 

import (
	"fmt"
	"reflect"
)

func example1() {
	var x1 int = 132
	x2 := 321
	x3 := "ewrdgdfgdfg"


	var typesOfVars = []reflect.Type {
		reflect.TypeOf(x1),
		reflect.TypeOf(x2),
		reflect.TypeOf(x3),
	}
	for _, typeOfVar := range typesOfVars {
		fmt.Printf(
			"type: %s, size = %d, comparable = %t\n",
			typeOfVar.String(),
			typeOfVar.Size(),
			typeOfVar.Comparable(),
		)
	}
}