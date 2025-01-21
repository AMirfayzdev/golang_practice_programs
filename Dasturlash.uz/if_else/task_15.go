package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	fmt.Scan(&a, &b, &c)
	if a > 0 && b > 0 && c > 0 && (a+b+c) == 180 {
		fmt.Println("true")
	} else {
		println(false)
	}

}
