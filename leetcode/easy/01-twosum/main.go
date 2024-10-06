package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, v := range nums {
		if val, ok := mp[v]; ok {
			return []int{val, i}
		}

		mp[target-v] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // {0, 1}
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // {1, 2}
	fmt.Println(twoSum([]int{3, 3}, 6))         // {0, 1}
}
