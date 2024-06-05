/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		_, exist := m[num-1]
		if exist {
			m[num-1] = num
		}
		m[num] = num
		_, exist = m[num+1]
		if exist {
			m[num] = num + 1
		}
	}
	maxLen := 0
	for _, num := range nums {
		_, exist := m[num-1]
		if exist {
			continue
		}
		maxLen = max(maxLen, getLen(num, &m))
	}
	return maxLen
}

func getLen(num int, m *map[int]int) int {
	if (*m)[num] == num {
		return 1
	}
	ret := 1 + getLen((*m)[num], m)
	(*m)[num] = num
	return ret
}

// @lc code=end
