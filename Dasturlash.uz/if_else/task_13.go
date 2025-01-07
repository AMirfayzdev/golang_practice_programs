package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if len(a) == 1 && a[0] >= '0' && a[0] <= '9' {
		fmt.Println("none")
	} else if len(a) == 1 && a[0] >= 'A' && a[0] <= 'Z' {
		fmt.Println("upperCase")
	} else if len(a) == 1 && a[0] >= 'a' && a[0] <= 'z' {
		fmt.Println("lowerCase")
	}
}
