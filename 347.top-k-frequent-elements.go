/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
import "sort"

type Pair struct {
	key   int
	value int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, value := range nums {
		_, exist := m[value]
		if exist == true {
			m[value]++
		} else {
			m[value] = 1
		}
	}
	temp := make([]Pair, 0)
	for key, value := range m {
		temp = append(temp, Pair{key, value})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].value > temp[j].value
	})
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, temp[i].key)
	}
	return res
}

// @lc code=end
