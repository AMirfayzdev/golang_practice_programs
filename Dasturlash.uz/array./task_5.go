package main

import "fmt"

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}
	max := 0
	fmt.Scan(&max)
	if max < 0 || max >= len(nums) {
		fmt.Println(0)
	} else {
		fmt.Println(nums[max])
	}
}
