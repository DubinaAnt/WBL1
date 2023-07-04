package main

/*Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20. */

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(4230483441509355), big.NewInt(29485204934792) //создаем большие числа

	mul := new(big.Int).Mul(a, b) //втстроенное в пакет умнножениe
	div := new(big.Int).Div(a, b) //втстроенное в пакет делениe
	sum := new(big.Int).Add(a, b) //втстроенное в пакет сложениe
	sub := new(big.Int).Sub(a, b) //втстроенное в пакет вычитаниe

	fmt.Printf("а : %d b : %d \n", a, b)
	fmt.Println("Результат умножения:", mul)
	fmt.Println("Результат деления:", div)
	fmt.Println("Результат сложения:", sum)
	fmt.Println("Результат вычитания:", sub)
}
