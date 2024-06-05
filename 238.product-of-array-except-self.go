/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	var allMul int64 = 1
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			allMul *= int64(nums[i])
		} else {
			zeroCount++
			if zeroCount > 1 {
				return make([]int, len(nums))
			}
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if zeroCount > 0 {
				res = append(res, 0)
			} else {
				res = append(res, int(allMul/int64(nums[i])))
			}
		} else {
			res = append(res, int(allMul))
		}
	}
	return res
}

// @lc code=end

