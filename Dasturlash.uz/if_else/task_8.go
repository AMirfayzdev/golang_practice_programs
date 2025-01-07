package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("musbat:", a)

	} else if a < 0 {
		fmt.Println("manfiy:", a)

	} else {
		fmt.Println(a)

	}
}
