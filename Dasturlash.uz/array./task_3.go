package main

import "fmt"

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		max += nums[i]

	}

	fmt.Println(max / len(nums))

}
