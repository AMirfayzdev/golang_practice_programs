// a va b sonlar orasida 3ga bo'linadiganlar

// Misol Sharti

// a, b sonlar berilgan.

// a va b sonlar orasidagi 3 ga bo'linadigan sonlarni yig'indisi va sonini konsolga chiqaring. 



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
		if i%3 == 0 {
			count += i

		}

	}
	fmt.Println("yig'indi = ", count)
	fmt.Println("soni = 3")
}
