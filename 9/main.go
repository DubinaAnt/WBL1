package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	chanel1 := chanel1(array)
	chanel2 := chanel2(chanel1)

	for item := range chanel2 {
		fmt.Println(item)
	}
}

func chanel1(array []int) <-chan int {
	chanel := make(chan int) //создаем канал
	go func() {              //запускаем горутину
		for _, item := range array { //пишем все из массива в канал
			chanel <- item
		}
		close(chanel) //закрываем тогда когда все значения запишутся
		//будет записывать по одному потому что нет буфера
	}()
	return chanel //возвращаем закрытый канал из которого мы сможем только читать послднее значение оставшееся там
}

func chanel2(chanel1 <-chan int) <-chan int {
	chanel := make(chan int)
	go func() {
		for item := range chanel1 {
			chanel <- item * 2
		}
		close(chanel)
	}()
	return chanel
}
