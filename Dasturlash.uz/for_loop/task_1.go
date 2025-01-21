// Sonlar yig'indisi

// Misol Sharti
// a soni berilgan.  0 dan a gacha bo'lgan sonlaryig'indisini toping  va yig'indini konsolga chiqaring. a sonini ham hisobga oling.

package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	count := 0
	for i := 0; i <= a; i++ {
		count += i
	}
fmt.Println(count)
}
