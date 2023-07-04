package main

/*Разработать программу, которая будет последовательно
отправлять значения в канал, а с другой стороны канала — читать.
 По истечению N секунд программа должна завершаться. */

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	timeSecond := 0
	fmt.Println("Укажите количество секунд до завершения:")
	_, err := fmt.Scan(&timeSecond)
	if err != nil {
		log.Fatal("Вы ввели не число")
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(timeSecond))
	//функция WithTimeout в качестве входных данных используется длительность времени
	//эта функция возвращает контекст, который отменяется при истечении времени
	//time.Duration - время прошедшее между двумя моментами, в виде наносекунд int64
	var wg sync.WaitGroup             //создаем группу для наших горутин
	channel := make(chan interface{}) //канал для чтения и записи

	wg.Add(2) //добавляем в группу 2 горутины
	go publisher(channel, ctx, &wg)
	go subscriber(channel, ctx, &wg)
	wg.Wait() //как только счетчик в который мы добавили в Add станет 0, то программа продолжит(завершит) выполняться
}

func publisher(pupChannel chan<- interface{}, ctx context.Context, wg *sync.WaitGroup) { //кидаем канал только для записи
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): //если в канал поступает уведомление о завершении работы программы
			wg.Done() //уменьшаем счетчик в группе на 1
			return    //выходим из цикла
		default: //если case не выполняется , то выполняется дефолт
			if i%2 == 0 {
				pupChannel <- fmt.Sprintf("text %d", i) //кидаем данные в канал
				time.Sleep(2 * time.Second)
				continue
			}
			pupChannel <- i //тоже кидаем в канал
			time.Sleep(2 * time.Second)
		}
	}
}

func subscriber(subscriber <-chan interface{}, ctx context.Context, wg *sync.WaitGroup) { //передаем канал только на чтение
	for {
		select {
		case <-ctx.Done(): //если в канал поступает уведомление о завершении работы программы
			wg.Done() //уменьшаем счетчик в группе на 1
			return    //выходим из цикла
		case item := <-subscriber: //если в канал поступают данные , то пишем их в item
			fmt.Println(item) //выводим данные
		} //цикл продолжается дальше по новой
	}
}
