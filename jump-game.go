package dsalgo

/*
Sort of dp problem we need to realise that we should start from the end
once we start from last position should n-1 + jump[i]>=n
basically if we jump from previous position will we be able to reach next index
if yes we update latest_idx
Now when there is the case of 0 what we do is we cant jump so latest_idx is the same
now at the end if we reach 0 index it means we can jump there
*/
func canJump(nums []int) bool {
	last_idx := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= last_idx {
			last_idx = i
		}
	}
	return last_idx == 0
}
