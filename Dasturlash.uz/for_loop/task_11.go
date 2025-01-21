// a sonini 5chi darajasi

// Misol Sharti

// a soni berilgan.

// a sonini 5 chi darajasini hisoblab konsolga chiqaring

package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	c := 1
	for i := 0; i < 5; i++ {
		c = c * a

	}
	{
		fmt.Println("Natija:", c)
	}

}
