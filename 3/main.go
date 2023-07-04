package main

/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений*/

import (
	"fmt"
)

func main() {
	var arr []int = []int{2, 4, 6, 8, 10}
	channel := make(chan int)
	fmt.Println(arr)
	for _, a := range arr { //считаем квадраты и отправляем в канал
		go Square(a, channel)
	}
	var sum int
	for i := 0; i < len(arr); i++ { //суммируем значения из канала
		sum += <-channel
	}

	fmt.Println(sum)
}

func Square(num int, channel chan int) {
	channel <- num * num

}
