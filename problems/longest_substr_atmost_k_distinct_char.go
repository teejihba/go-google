package problems

func LongestSubstringWithAtmostKDisntinctChar(str string, k int) (int, int) {
	table := make(map[uint8]int)
	start := 0
	end := 0
	ans := 0
	index := 0
	total := 0
	for end < len(str) {
		if val, ok := table[str[end]]; ok {
			table[str[end]] = val + 1
			end++
			total++
			if total > ans {
				ans = total
				index = start
			}
		} else {
			if len(table) < k {
				table[str[end]] = 1
				total++
				if total > ans {
					ans = total
					index = start
				}
				end++
			} else {
				for len(table) == k {
					if val_, _ := table[str[start]]; val_ > 1 {
						table[str[start]] -= 1
					} else {
						delete(table, str[start])
					}
					total--
					start++
				}
			}
		}
	}
	return ans, index
}
