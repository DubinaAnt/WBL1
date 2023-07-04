package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctxCansel, cancel := context.WithCancel(context.Background()) //начинает завершение когда вызыван метод cancel
	//cansel не рекомендуют передавать в функции, надо вызывать там где был обьявлен котнекст
	ctxDeadline, _ := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
	//Эта функция возвращает производный контекст от своего родителя, который отменяется после дедлайна
	//или вызова функции отмены. Например, вы можете создать контекст, который автоматически отменяется в
	//определенное время
	ctxTimeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//Эта функция похожа на context.WithDeadline. Разница в том, что в качестве входных данных
	//используется длительность времени. Эта функция возвращает производный контекст, который отменяется
	//при вызове функции отмены или по истечении времени.

	done := make(chan os.Signal, 1) //через сигнал , как мы уже делали в 4 задании
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(4)

	go signalChanel(done, &wg)
	go ctxWithCancel(ctxCansel, &wg)
	go ctxWithDeadline(ctxDeadline, &wg)
	go ctxWithTimeout(ctxTimeout, &wg)

	time.Sleep(5 * time.Second)
	cancel() //завершает нашу первую горутину с помощью WithCancel контекста

	wg.Wait()

}

func ctxWithCancel(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel")
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("cancel func")
		}
	}
}

func ctxWithDeadline(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("deadline")
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("deadline func")
		}
	}
}

func ctxWithTimeout(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("timeout func")
		}
	}
}

func signalChanel(cancelChan <-chan os.Signal, wg *sync.WaitGroup) {
	for {
		select {
		case <-cancelChan:
			fmt.Println("cancelChan")
			wg.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("cancelChan func")
		}
	}
}
