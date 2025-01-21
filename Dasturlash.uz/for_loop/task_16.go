package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanln(&a)
	count := 0
	for i := a; i > 0; i = i / 10 {
		count++
	}

	fmt.Println(count)

}
