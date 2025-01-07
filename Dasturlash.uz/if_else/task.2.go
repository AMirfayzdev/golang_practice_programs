package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 10 {
		fmt.Println(n + 3)
	} else if n < 10 {
		fmt.Println(n * 2)
	} else if n == n {
		fmt.Println(22)
	}
}