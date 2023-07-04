package main

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun rice"
	fmt.Println(wordFlip(text))
}

func wordFlip(text string) string {
	arrWord := strings.Fields(text)                          //разбивает строку по отдельным словам и возвращает слайс слов
	for i, j := 0, len(arrWord)-1; j >= i; i, j = i+1, j-1 { //просто бежим двумя индексами i j навстречу друг другу
		arrWord[i], arrWord[j] = arrWord[j], arrWord[i] //и меняем местами элементы
	}
	return strings.Join(arrWord, " ") //а тут объединяет слайс слов в строку с разделителем " "

}
