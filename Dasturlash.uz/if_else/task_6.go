package main

import "fmt"

func main() {
	a := 0
	b := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		fmt.Println("katta son:", a)

	} else if a < b {
		fmt.Println("katta son:", b)

	} else if a == b {
		fmt.Println("Voy ular teng")

	}
}
