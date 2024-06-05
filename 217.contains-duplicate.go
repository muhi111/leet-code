/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, exists := m[nums[i]]
		if exists == true {
			return true
		} else {
			m[nums[i]] = 1
		}
	}
	return false
}

// @lc code=end

