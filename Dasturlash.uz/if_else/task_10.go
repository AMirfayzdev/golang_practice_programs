package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	if a%3 == 0 && a%4 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
