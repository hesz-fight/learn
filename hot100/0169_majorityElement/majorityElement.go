package majorityelement

func majorityElement(nums []int) int {
	var votes, cand int
	for _, num := range nums {
		if votes == 0 {
			cand = num
			votes++
			continue
		}
		if num != cand {
			votes--
		} else {
			votes++
		}
	}
	return cand
}
