package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	var p int
	var maxCount int

	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {

		for {

			if _, ok := m[s[i]]; ok {
				delete(m, s[p])
				p++
			} else {
				break
			}

		}

		m[s[i]] = true

		maxCount = max(maxCount, i-p+1)

	}

	return maxCount
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
