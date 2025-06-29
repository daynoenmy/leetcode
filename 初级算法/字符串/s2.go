package str

func firstUniqChar(s string) int {
	pos := [26]int{}
	n := len(s)
	for i := range pos {
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	for _, p := range pos {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}
