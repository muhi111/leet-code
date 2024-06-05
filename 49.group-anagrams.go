/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]int)
	// O(n*mlogm)
	for i, value := range strs {
		sortedStr := SortString(value)
		_, exist := m[sortedStr]
		if exist == true {
			m[sortedStr] = append(m[sortedStr], i)
		} else {
			m[sortedStr] = []int{i}
		}
	}
	// O(n*n)
	for _, value := range m {
		temp := []string{}
		for _, index := range value {
			temp = append(temp, strs[index])
		}
		res = append(res, temp)
	}
	return res
}

// O(mlogm)
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, exist := m[s[i]]
		if exist == true {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
		_, exist = m[t[i]]@
		if exist == true {
			m[t[i]]--
		} else {
			m[t[i]] = -1
		}
	}
	for value := range m {
		if m[value] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
