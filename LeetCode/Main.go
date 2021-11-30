package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for i, elem := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (elem + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
