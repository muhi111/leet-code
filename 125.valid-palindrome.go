/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
import "strings"

func isPalindrome(s string) bool {
	t := ""
	for i := 0; i < len(s); i++ {
		if isAlphaNum(s[i]) {
			t += string(s[i])
		}
	}
	t = strings.ToLower(t)
	for i := 0; i < len(t)/2; i++ {
		if t[i] != t[len(t)-1-i] {
			return false
		}
	}
	return true
}

func isAlphaNum(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

// @lc code=end
