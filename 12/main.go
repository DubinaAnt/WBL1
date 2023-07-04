package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	res := uniqItems(input)

	fmt.Println(res)
}

func uniqItems(input []string) []string {
	uniqItemsMap := map[string]struct{}{} //ключи в map уникальны, пустая структура ничего не "весит"

	for _, item := range input {

		uniqItemsMap[item] = struct{}{}
	}

	var res []string

	for key := range uniqItemsMap { //добавляем ключи в слайс, значения не нужны они и так пустые
		res = append(res, key)
	}
	return res
}
