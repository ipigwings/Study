package main

import "fmt"

func main() {
	var test []int = []int{1, 2, 3, 1, 1, 3}
	fmt.Print(numIdenticalPairs(test))

}

func numIdenticalPairs(nums []int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return res
}
