package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, v := range nums {
		y := target - v

		if x, ok := m[y]; ok {
			return []int{idx, x}
		}
		m[v] = idx
	}

	return []int{-1, -1}
}
