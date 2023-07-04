package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10) //создаем большую строку
	justString = v[:100]           //если сделать слайс таким образом он не обрежет строку по 100ый символ, а выдаст символы с 1 по 100 индекс, и последний элемент может потерятся
	// в таком случаее лучге пройтись for range
	var arr []string
	for index, value := range v {
		arr = append(arr, value)
	}
	justString = strings.Join(arr, "")
}

func main() {
	someFunc()
}
