package main

/*Разработать программу, которая переворачивает подаваемую
на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.*/

import "fmt"

func main() {
	text := "T日本語"
	fmt.Println(Reverse(text))
}

func Reverse(text string) string {
	arrRune := []rune(text)                                  //преобразуем строку в слайс рун
	for i, j := 0, len(arrRune)-1; j >= i; i, j = i+1, j-1 { //проходим циклом слайс с двух сторон и меняем местами
		arrRune[i], arrRune[j] = arrRune[j], arrRune[i]
	}
	return string(arrRune) //преобразуем в строку
}
