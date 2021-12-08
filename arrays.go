package arrays

import "fmt"

func BinarySearch(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return true
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

func welcome() {
	fmt.Println("Welcome to array")
}
func Welcome() {
	fmt.Println("This package was created by Rizki Susilo Ramadhan")
}
