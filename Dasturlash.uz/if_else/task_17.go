package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	fmt.Scan(&a, &b, &c)

	if a == b && c == a && b == c {
		println("equilateral")
	} else if a == c {
		println("isosceles")

	} else {
		println("scalene")
	}
}
