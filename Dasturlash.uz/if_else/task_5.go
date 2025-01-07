package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("Juft")
	} else if a%2 > 0 {
		fmt.Println("Toq")
	}
}
