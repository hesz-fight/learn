package majorityelem

func majorityelem(nums []int) int {
	ans := 0
	set := map[int]int{}
	for _, num := range nums {
		set[num] += 1
		if set[num] > len(nums)/2 {
			ans = num
			break
		}
	}
	return ans
}

func majorityelem2(nums []int) int {

	votes, canidate := 0, -1
	for _, num := range nums {
		if votes == 0 {
			votes++
			canidate = num
			continue
		}
		if num == canidate {
			votes++
		} else {
			votes--
		}
	}
	return canidate
}
