/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
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
		_, exist = m[t[i]]
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

