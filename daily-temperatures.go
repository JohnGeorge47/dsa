package dsalgo

// https://leetcode.com/problems/daily-temperatures/
type stackItem struct {
	pos  int
	temp int
}

/*
What we first do here is move from right to left  and also we maintian a  stack
whenever we encounter an element we do a check
if element < top of stack
      if yes we push element to stack and find the diff between positions
if element >= top of stack
      now in this case it means we have to discard top of stack element because this wont be closest greater
      temp for any element before current element so we keep ejecting elements from stack
      until we reach len(stack)=0 or top of stack is greater
During every iteration we add element to stack
*/

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]stackItem, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		result[i] = 0
		for len(stack) != 0 && temperatures[i] >= stack[len(stack)-1].temp {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1].pos - i
		}
		stack = append(stack, stackItem{pos: i, temp: temperatures[i]})
	}
	return result
}
