package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type User struct {
	Person
	Email string
}
type Employee struct {
	JobTitle string
}
type User_add struct {
	Person
	Employee
	Email string // очень похоже на множественное наследование, но это не совсем оно
}
type Student struct {
	User   User
	Course struct { //можно описать структуру внутри структуры, но так лучше не делать
		Name string
	}
}
func ( Person) Say() string {
	return "Hello"
}
func (p Person) Oldify() {
	p.Age +=100
}
func (p *Person) OldifyPtr() {
	p.Age +=100
}

// описание методов для структуры
func (u User) String() string { //"u User" - ресивер(то, для чего хотим сделать метод)(тут "u" - просто локальная переменная), "String()"-название метода
	return fmt.Sprintf("User Name: %s, Email: %s", u.Name, u.Email)
}







func main() {
	p := Person{
		Name: "Alice",
		Age: 30,
	}
	p.Oldify()
	fmt.Println(p.Age) //30, тк передаем значение и не возвращаем ничего
	p.OldifyPtr()
	fmt.Println(p.Age) //130


	var person1 = Person{
		Name: "Loki",
		Age:  20, //в этом случае необязательно указывать все атрибуты структуры
	}
	var _ = Person{
		"Bob",
		20, //обязательно указывать все атрибуты в нужном порядке, иначе CE(лучше избегать)
	}
	var person = User{
		Person: Person{ //когда используем структуру с вложенной структурой делаем так
			Name: "Bob",
			Age:  34,
		},
		Email: "example@mail.ru",
	}
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(User.String(person))
	fmt.Println(person.String()) //тоже самое что и строкой выше(вот так лучше писать)
	fmt.Println(person1)
	fmt.Println(person.Say()) //Say() принадлежит struct Person, которая вложена в struct User, поэтому и работает
}
