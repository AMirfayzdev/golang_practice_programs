// a sonini nchi darajasi

// Misol Sharti

// a , n sonlari berilgan (a soni 1-9 o'rtasida ).

// Ketma ketlikni hisoblang. a + aa + aaa + ........  nta a bo'ladi.

// Yig'indini konsolga chiqarish kerak.

package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	fmt.Scan(&a, &b)
	count := 0
	ar := 0
	for i := 0; i < b; i++ {
		count = count*10 + a
		ar += count
		fmt.Println(count)
		fmt.Println(ar)
	}
	{
		fmt.Println(ar)
	}
}
