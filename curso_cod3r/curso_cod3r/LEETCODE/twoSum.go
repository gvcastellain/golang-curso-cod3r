package main

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums)-1; i++ {
		for j := 0; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

	// You may assume that each input would have exactly one solution, and you may not use the same element twice.

	// You can return the answer in any order.
}
