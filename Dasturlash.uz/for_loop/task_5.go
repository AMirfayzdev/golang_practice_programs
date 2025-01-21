// Juft sonlar soni

// Misol Sharti


// a soni berilgan. 0 dan a gacha bo'lgan juft sonlar soni toping va konsolga chiqaring. a sonini inobatga olmaymiz.

package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	count := 0
	for i := 0; i < a; i++ {
		if i%2 == 0 {
			count ++ 
		}
	}
	fmt.Println(count)
}
