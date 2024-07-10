/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// numsをソートする
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 2分探索でtargetを探す
			target := -nums[i] - nums[j]
			left, right := j+1, len(nums)-1
			for left <= right {
				mid := (left + right) / 2
				if nums[mid] == target {
					result = append(result, []int{nums[i], nums[j], nums[mid]})
					break
				} else if nums[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}

	if len(result) == 0 {
		return result
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i][0] < result[j][0] {
			return true
		} else if result[i][0] > result[j][0] {
			return false
		}
		if result[i][1] < result[j][1] {
			return true
		} else if result[i][1] > result[j][1] {
			return false
		}
		if result[i][2] < result[j][2] {
			return true
		} else {
			return false
		}
	})

	unique_result := make([][]int, 0)
	unique_result = append(unique_result, result[0])

	for i := 1; i < len(result); i++ {
		if result[i-1][0] == result[i][0] && result[i-1][1] == result[i][1] && result[i-1][2] == result[i][2] {
			continue
		}
		unique_result = append(unique_result, result[i])
	}

	return unique_result
}

// @lc code=end
