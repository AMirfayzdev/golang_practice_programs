package main

import "fmt"

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}
	max := make([]int, len(nums) +n) 
	n := 0 

	fmt.Scan(&n)
	for i :=0; i < len(max); i++{
        if i < len(max)-1{
        max[i] = nums[i]
    }else{
          max[i] = n
    }
    }

  fmt.Println(max)
}