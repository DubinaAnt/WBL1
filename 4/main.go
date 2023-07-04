package main

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные
данные из канала и выводят в stdout. Необходима возможность выбора
количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. */

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var workerCount int
	fmt.Println("Укажите количество воркеров:")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		log.Fatal("Вы ввели не число")
	}

	channel := make(chan interface{})                    //канал для воркеров который принимает любые данные
	done := make(chan os.Signal, 1)                      //буферезированный канал который принимает один сигнал от ОС
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM) //тут мы ловим и записываем сигнал от ОС в канал

	for i := 0; i < workerCount; i++ {
		go worker(channel) //запускам столько воркеров сколько указал пользователь
	}

	for i := 0; ; i++ {
		select { // оператор итерации по каналам. причем будет исполнен каждый кейс в отличии от свича
		case <-done: //читаем канал done и если туда что-то приходит то
			close(channel) //закрываем канал для воркеров
			close(done)    //закрываем канал для чтение сигналов ос
			return         //выходим из бесконечного цикла и основаня горутина мейн тоже завершится
		default:
			if i%2 == 0 {
				channel <- fmt.Sprintf("text %d", i) //пишем в канал воркеров
				time.Sleep(2 * time.Second)
				continue
			}
			channel <- i // тоже пишем
			time.Sleep(2 * time.Second)
		}
	}

}

func worker(chanel <-chan interface{}) { //принимает канал доступный тут только для чтения
	for item := range chanel { //циклом по каналу доставая все что сюда приходит,
		fmt.Println(item) //  цикл будет бесконечно читать из канала пока он не закроется и потом
	} //выход из цикла и горутина завершится
}
