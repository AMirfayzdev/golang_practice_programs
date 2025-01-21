// Toq sonlar yig'indisi

// Misol Sharti

// a soni berilgan. 0 dan  a gacha bo'lgan toq sonlar yig'indisini toping  va yig'indini konsolga chiqaring. a sonini ham inobatga oling.



package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	count := 0
	for i := 0; i <= a; i++ {
		if i%2 == 1 {
			count += i
		}
	}
	fmt.Println(count)
}
