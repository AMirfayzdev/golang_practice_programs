package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	if a%5 == 0 {
		fmt.Println("true")
	} else if a%5 > 0 {
		fmt.Println("false")
	}
}
