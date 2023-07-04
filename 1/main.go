package main

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).*/

import "fmt"

type Human struct {
	name   string
	age    int
	gender string
}

func (h Human) GetName() {
	fmt.Println("Имя", h.name)
}

func (h Human) GetAge() {
	fmt.Println("Возраст", h.age)
}

func (h Human) GetGender() {
	fmt.Println("Пол", h.gender)
}

type Action struct {
	Human
}

func main() {
	var action = Action{Human{"Иван", 28, "Мужской"}}
	action.GetName()
	action.GetAge()
	action.GetGender()
}
