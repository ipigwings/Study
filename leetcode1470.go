package main

import "fmt"

func main() {
	test1 := []int{2, 5, 1, 3, 4, 7}
	fmt.Print(shuffle(test1, 3))
}

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	var j int
	for i := range nums {
		if i >= n {
			break
		}
		res[j] = nums[i]
		res[j+1] = nums[i+n]
		j += 2
	}
	return res
}
