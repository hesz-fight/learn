package productexceptself

func productExceptSelf(nums []int) []int {
	// 前缀乘积
	accu := 1
	prefix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		accu *= nums[i-1]
		prefix[i] = accu
	}
	// 后缀乘积
	accu = 1
	suffix := make([]int, len(nums))
	suffix[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		accu *= nums[i+1]
		suffix[i] = accu
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = suffix[i] * prefix[i]
	}

	return ans
}
