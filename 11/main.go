package main

// Реализовать пересечение двух неупорядоченных множеств

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{4, 5, 0}

	сross := crossing(array1, array2)
	fmt.Println(сross)
}

func crossing(array1, array2 []int) []int {
	inMap := map[int]struct{}{} //в map ключи уникальные

	for _, item1 := range array1 { //бежим каждым элементом первого массива
		for _, item2 := range array2 { //по каждому элементу второго массива
			if item1 == item2 { //если значение из первого массива равно значению второго
				inMap[item1] = struct{}{} ///если в одном из масивов несколько одиноковых элементов , то в мапу перезапишется ключ
			}
		}
	}

	var cross []int

	for itemMap := range inMap { //просто переписываем из ключей мапы значения в слайс
		cross = append(cross, itemMap)
	}
	return cross //возвращаем слайс
}
