// https://leetcode.com/problems/search-insert-position
func searchInsert(nums []int, target int) int {
	idx := 0

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	// variable to hold value that higher than target, but lowest from so on
	// eg := 1,5,6,8 > target = 4. highest should 5 and the index is 1

	highestValue := target
	firstFind := false
	for index, num := range nums {
		if target == num {
			idx = index
			break
		}

		if !firstFind && num > highestValue {
			firstFind = true
			idx = index
		}
	}
	return idx
}