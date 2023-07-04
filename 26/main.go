package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "avc"

	for _, str := range input {
		if strings.Count(input, string(str)) > 1 {
			fmt.Printf("%v - false", input)
			break
		} else {
			fmt.Printf("%v - true", input)
			break
		}
	}

}
