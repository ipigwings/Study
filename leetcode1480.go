package main

import "fmt"

func main() {
	test1 := []int{1, 2, 3, 4}
	fmt.Print(runningSum(test1))
}

func runningSum(nums []int) []int {
	res := make([]int, len(nums))

	var s int
	for i, v := range nums {
		s += v
		res[i] = s
	}

	return res
}
