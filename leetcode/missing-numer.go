package leetcode

func missingNumber(nums []int) int {
	numsSums := 0
	sums := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		numsSums = numsSums + nums[i]
		sums = sums + i
	}
	sums = sums + n
	return sums - numsSums
}
