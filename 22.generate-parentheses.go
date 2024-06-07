/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start

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

func generateParenthesis(n int) []string {
	res := []string{}
	stack := Stack{}
	generateParenthesisRecursive(2*n, &stack, &res, "")
	return res
}

func generateParenthesisRecursive(n int, stack *Stack, res *[]string, current string) {
	if n == 0 {
		if len(*stack) == 0 {
			*res = append(*res, current)
		}
		return
	}
	stack.Push("(")
	generateParenthesisRecursive(n-1, stack, res, current+"(")
	stack.Pop()
	if stack.Pop() == "(" {
		generateParenthesisRecursive(n-1, stack, res, current+")")
		stack.Push("(")
	}
}

// @lc code=end
