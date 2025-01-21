// Juft sonlar yig'indisi

// Misol Sharti

// a soni berilgan. 0 dan a gacha bo'lgan juft  sonlar yig'indisini toping  va yig'indini konsolga chiqaring

package main

import (
	"fmt"
)

func main() {
	a := 10
	count := 0
	for i := 0; i <= a; i++ {
		if (i % 2 == 0){
		count += i
	}
	}
fmt.Println(count)
}