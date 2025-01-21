// a va b sonlar orasidagi sonlar yig'indisi

// Misol Sharti

// a, b sonlar berilgan.

// a va b sonlar orasidagi (butun) sonlar   yig'indisini toping (a sonini inobatga oling) va yig'indini konsolga chiqaring.




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
		{
			count += i
		}
	}
	fmt.Println(count)
}
