package main

// Дана переменная int64. Разработать программу
// которая устанавливает i-й бит в 1 или 0.

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64 = 57678
	var pos uint = 5
	fmt.Println(strconv.FormatInt(num, 2))
	if hasBit(num, pos) == true {
		num = clearBit(num, pos)
		fmt.Println(strconv.FormatInt(num, 2))
		fmt.Println(num)
	} else {
		num = setBit(num, pos)
		fmt.Println(strconv.FormatInt(num, 2)) //перевод в 2чную систему
		fmt.Println(num)
	}
	// fmt.Println(hasBit(num, pos))
	// num = clearBit(num, pos)
	// fmt.Println(hasBit(num, pos))
	// fmt.Println(strconv.FormatInt(num, 2))
	// fmt.Println(hasBit(num, pos))
}

// ставит бит
func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

// удаляет бит
func clearBit(n int64, pos uint) int64 {
	var mask int64
	mask = ^(1 << pos)
	n &= mask
	return n
}

// Получение значения бита
func hasBit(n int64, pos uint) bool {
	var val int64
	val = n & (1 << pos)
	return (val > 0)
}
