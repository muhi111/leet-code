/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start

// stackを実装する
type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func isValid(s string) bool {
	stack := Stack{}
	for _, c := range s {
		str := string(c)
		if str == "(" || str == "[" || str == "{" {
			stack.Push(str)
		} else if str == ")" || str == "]" || str == "}" {
			top := stack.Pop()
			if (str == ")" && top != "(") || (str == "]" && top != "[") || (str == "}" && top != "{") {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end
