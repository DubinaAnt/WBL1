package main

/* Разработать программу, которая в рантайме способна
определить тип переменной: int, string, bool, channel
из переменной типа interface{} */

import "fmt"

func WhatType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case chan int:
		fmt.Println("chan")
	case []int:
		fmt.Println("slice")
	case bool:
		fmt.Println("bool")
	case float64:
		fmt.Println("slice")
	default:
		fmt.Printf("unknown")
	}
}

func main() {
	channel := make(chan int)
	slice := make([]int, 1, 1)
	WhatType(21)
	WhatType("hello")
	WhatType(32.4523)
	WhatType(channel)
	WhatType(true)
	WhatType(slice)
	WhatType([4]int{})
}
