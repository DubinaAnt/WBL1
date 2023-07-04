package main

// удалить i элемент из слайса

import (
	"fmt"
)

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	b := []string{"A", "B", "C", "D", "E"}
	i := 1
	a = fastCut(a, i)
	fmt.Println(a, len(a), cap(a))
	b = slowCut(b, i)
	fmt.Println(b, len(b), cap(b))

}

func fastCut(slice []string, i int) []string {
	// 1. Копировать последний элемент в индекс i.
	slice[i] = slice[len(slice)-1]

	// 2. Удалить последний элемент (записать нулевое значение).
	slice[len(slice)-1] = ""

	// 3. Усечь срез.
	slice = slice[:len(slice)-1]
	return slice
}

func slowCut(slice []string, i int) []string {
	// 1. Выполнить сдвиг slice[i+1:] влево на один индекс.
	copy(slice[i:], slice[i+1:])

	// 2. Удалить последний элемент (записать нулевое значение).
	slice[len(slice)-1] = ""

	// 3. Усечь срез.
	slice = slice[:len(slice)-1]
	return slice
}
