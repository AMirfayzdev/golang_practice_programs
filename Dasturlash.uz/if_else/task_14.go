package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	d := a % 7
	if d == 1 {
		fmt.Println("Dushanba")
	} else if d == 2 {

		fmt.Println("Seshanba")
	} else if d == 3 {

		fmt.Println("Chorshanba")
	} else if d == 4 {
		fmt.Println("Payshanba")
	} else if d == 5 {

		fmt.Println("Juma")
	} else if d == 6 {
		fmt.Println("Shanba")
	} else if d == 0 {
		fmt.Println("Yakshanba")
	} else {
		fmt.Println("none")
	}
}
