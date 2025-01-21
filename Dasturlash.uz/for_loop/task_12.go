// a sonini nchi darajasi

// Misol Sharti

// a soni berilgan.

// a sonini n chi darajasini hisoblaydigan dastur yozing  va hosil bo'lgan sonni konsolga chiqaring.

package main

import (
	"fmt"
)

func main() {
	a := 0

	n := 0

	c := 1

	fmt.Scan(&a, &n)
	for i := 0; i < n; i++ {
		c = c * a
	}
	{
		fmt.Println("Natija:", c)
	}

}
