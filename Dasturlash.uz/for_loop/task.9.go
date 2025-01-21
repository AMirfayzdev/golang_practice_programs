// a va b sonlar orasida 2 va 3 ga bo'linadiganlar

// Misol Sharti

// a, b sonlar berilgan.

// a va b sonlar orasidagi 2 va 3 ga bo'linadigan sonlarning sonini  konsolga chiqaring.

package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	fmt.Scan(&a, &b)
	count := 0
	c := 0
	for i := a; i < b; i++ {
		if i%3 == 0 && i%2 == 0 {
			count++
			c += i

		}

	}
	fmt.Println("Natija:",count)

}
