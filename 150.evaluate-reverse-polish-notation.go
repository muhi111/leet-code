/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
import "strconv"

type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func isOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func operate(a, b int, op string) int {
	switch op {
	case "+":
		return b + a
	case "-":
		return b - a
	case "*":
		return b * a
	case "/":
		return b / a
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := Stack{}
	for _, token := range tokens {
		if isOperator(token) {
			stack.Push(operate(stack.Pop(), stack.Pop(), token))
		} else {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}
	return stack.Pop()
}

// @lc code=end
