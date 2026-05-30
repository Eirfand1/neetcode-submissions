func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		diff := target - v
		if j, found := m[diff]; found && j != i {
			return []int{i, j}
		}
	}
	return []int{0, 0}
}