package main

// Поменять местами два числа без создания временной переменной.

import "fmt"

func main() {
	a, b := 1, 3
	fmt.Println(a, b)
	a, b = b, a // сначала считываются правая часть, потом идет присваивание
	fmt.Println(a, b)
}
