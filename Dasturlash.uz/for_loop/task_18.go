package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanln(&a)
	count := 0
	for a > 0 {
		i := a % 10
		count = count*10 + i
		a /= 10

	}
	fmt.Println(count)
}