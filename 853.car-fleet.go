/*
 * @lc app=leetcode id=853 lang=golang
 *
 * [853] Car Fleet
 */

// @lc code=start
import (
	"sort"
)

type Car struct {
	position float64
	speed    int
}

type Stack []float32

func (s *Stack) Push(v float32) {
	*s = append(*s, v)
}

func (s *Stack) Top() float32 {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = Car{float64(position[i]), speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})
	stack := Stack{}
	for i := 0; i < len(position); i++ {
		time := (float32(target) - float32(cars[i].position)) / float32(cars[i].speed)
		if stack.Top() != -1 && stack.Top() >= time {
			continue
		}
		stack.Push(time)
	}
	return len(stack)
}

// @lc code=end
