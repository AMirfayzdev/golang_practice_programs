package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a > b && a > c {
		fmt.Println("katta son:", a)

	} else if b > a && b > c {
		fmt.Println("katta son:", b)

	} else if c > a && c > b {
		fmt.Println("katta son:" , c)

	}
}