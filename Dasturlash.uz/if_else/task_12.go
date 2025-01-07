package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if len(a) == 1 && a[0] >= '0' && a[0] <= '9' {
		fmt.Println("son")
	} else {
		fmt.Println("alpha")
	}
}
