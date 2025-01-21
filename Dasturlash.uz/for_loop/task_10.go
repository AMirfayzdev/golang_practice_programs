// a va b sonlar orasida musbat sonlar

// Misol Sharti

// a, b sonlar berilgan.

// a va b sonlari orasidagi musbat sonlar sonini toping  va shuni konsolga chiqaring.  b sonini ham inobatga oling.

package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	fmt.Scan(&a, &b)
	count := 0
	for i := a; i <= b; i++ {
		if i > 0 {
			count++

		}

	}
	fmt.Println("Natija:", count)

}
