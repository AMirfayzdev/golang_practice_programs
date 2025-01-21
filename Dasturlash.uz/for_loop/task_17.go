// Raqamlar yig'indisi
// Misol Sharti
// n  soni berilgan.
// Shu sonni raqamlarini yig'indisini toping.

package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanln(&a)
	count := 0
	for a > 0 {
		count += a % 10
		a /= 10

	}
	fmt.Println(count)
}
