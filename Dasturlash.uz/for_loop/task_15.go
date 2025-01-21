// Tub son




package main

import (
	"fmt"
)

func main() {

	a := 0
	fmt.Scan(&a)
	count := 0
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
