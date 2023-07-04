package main

/*Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.*/

import (
	"fmt"
	"sync"
)

func main() {
	var arr []int = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	fmt.Println(arr)
	for _, a := range arr {
		wg.Add(1)
		go Square(a, &wg)
	}
	wg.Wait()
}

func Square(a int, wg *sync.WaitGroup) {
	fmt.Println(a * a)
	wg.Done()
}
