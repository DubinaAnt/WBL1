package main

// Реализовать собственную функцию sleep

import (
	"fmt"
	"time"
)

func main() {

	sleep(1)
	fmt.Println("Прошло 1 секунды")
	sleep(2)
	fmt.Println("Прошло 2 секунды")
	sleep(3)
	fmt.Println("Прошло 3 секунды")
}

func sleep(d time.Duration) {
	<-time.After(d * time.Second)
}
