package main

// 没计算过/能到达最后一个坐标点/不能到达最后一个坐标点
var memo map[int]int = map[int]int{}

func canJump1(nums []int) bool {
	return traceback(0, len(nums), nums)
}

func traceback(curPos int, endPos int, nums []int) bool {
	if curPos >= endPos {
		return true
	}
	if memo[curPos] != 0 {
		return memo[curPos] > 0
	}

	for i := 1; i < nums[curPos]; i++ {
		can := traceback(curPos+i, endPos, nums)
		if can {
			memo[curPos] = 1
			return true
		}
		memo[curPos] = -1
	}

	return false
}
