package main

import (

)
type error interface { //так представлен тип ошибки в пакете errors(как я понял)
	Error() string
}
func doSomething() (string, error) //error возвращается последним типом(так принято)
//если вернули ошибку, то другие типы данных возвращаются со значениями по умолчанию
