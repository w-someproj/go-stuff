package main

import "fmt"

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

// optimized
func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, elem := range nums {
		if elIndex, ok := sumMap[(target - elem)]; ok {
			return []int{elIndex, i}
		} else {
			sumMap[elem] = i
		}
	}
	return []int{}
}
