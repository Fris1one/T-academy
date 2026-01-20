package main

import (
	"log"
	"errors"
	"fmt"
)
//1 способ создания своего типа ошибки
var ErrNotFound = errors.New("not found") //так делать не оч, тк их можно сравнивать только через ==

func doSomething1() error {
	return ErrNotFound
}

//2 способ
func divide(a,b int) (int, error) { //лучше использовать благодаря форматированиям(особенно %w)
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a) //под капотом в Errorf происходит errors.New и передается строка в ()
	}
	return a/b, nil
}
//3 способ

type errCustomString string //создали свой тип ошибки на основе строки
func (e errCustomString) Error() string {
	return string(e)
}
var ErrCustomString errCustomString = "validation failed"

//4 способ - самый лучший

type ErrCustom struct {
	Msg string
	Code int
	Op string
}

func (e ErrCustom) Error() string {
	return fmt.Sprintf("[%d] %s during %s", e.Code, e.Msg, e.Op)
}

func NewErrCustom (msg string, code int, op string) error {
	return ErrCustom{Msg: msg, Code: code, Op: op}
}
func main() {
	err := doSomething1()
	if err != nil {
		log.Fatal(err)
	}
	
	err1 := NewErrCustom("access deenied", 403, "auth.Check")

}