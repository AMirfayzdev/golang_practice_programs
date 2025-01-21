package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanln(&a)
	b := a
	count := 0
	for a > 0 {
		i := a % 10
		count = count*10 + i
		a /= 10

	}
	if count == b {
		fmt.Println("true")

	} else {
		fmt.Println("false")
	}
}
