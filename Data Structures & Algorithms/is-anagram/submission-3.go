func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		m1[v]++
		m2[rune(t[i])]++
	}
	fmt.Println(m1)
	fmt.Println(m2)
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}