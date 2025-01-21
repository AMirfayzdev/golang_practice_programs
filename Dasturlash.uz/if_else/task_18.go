package main

import "fmt"

func main() {

	a := 0
	b := 0

	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println(a)
	} else if b < a {
		fmt.Println(b)
	} else {
		fmt.Println("teng")
	}
}
l