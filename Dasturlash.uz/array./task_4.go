package main

import "fmt"

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}
	max := 0
	fmt.Scan(&max)
	status := false

	for i := 0; i < len(nums); i++ {
		if max == nums[i] {
			status = true
		}

	}

	fmt.Println(status)

}
