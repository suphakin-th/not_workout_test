package oddint

func FindOddInt(arr []int) int {
	// LeetCode Again!!!!! why you guy love it!!
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}
	return 0
}
