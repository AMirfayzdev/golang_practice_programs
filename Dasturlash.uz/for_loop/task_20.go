package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanln(&a)
	count := 1
	for i := 1; i <= a; i++ {
		count *= i
		
	} 
	fmt.Println(count)

}
