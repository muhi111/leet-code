/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

func dailyTemperatures(temperatures []int) []int {
	stack := Stack{}
	res := make([]int, len(temperatures))
	for i, t := range temperatures {
		for stack.Top() != -1 && temperatures[stack.Top()] < t {
			res[stack.Top()] = i - stack.Top()
			stack.Pop()
		}
		stack.Push(i)
	}
	return res
}

// @lc code=end
