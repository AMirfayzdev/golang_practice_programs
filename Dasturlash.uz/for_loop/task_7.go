// a va b sonlar orasidagi juft sonlar

// Misol Sharti

// a, b sonlar berilgan
// a va b sonlar orasidagi juft sonlarni sonini toping
// (a sonini inobatga oling) va natijani konsolga chiqaring.


package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	fmt.Scan(&a, &b)
	count := 0
	for i := a; i < b; i++ {
		if i%2 == 0 {
			count++
		}

	}
	fmt.Println(count)
}
