package ArraysDo

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
