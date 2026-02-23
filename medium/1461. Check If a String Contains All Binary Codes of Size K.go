package medium

func hasAllCodes260223(s string, k int) bool {
	dict := map[string]bool{}
	for i := 0; i <= len(s)-k; i++ {
		dict[s[i:i+k]] = true
	}
	// fmt.Println(2<<(k-1), dict)
	return len(dict) == (2 << (k - 1))
}

func hasAllCodes2602232(s string, k int) bool {
	if len(s) <= k {
		return false
	}
	dict := make([]bool, 2<<(k-1))
	curr := 0
	for i := 0; i < k; i++ {
		curr = curr<<1 + int(s[i]-'0')
	}
	dict[curr] = true
	mod := 1 << k
	// fmt.Println(curr, mod)
	for i := k; i < len(s); i++ {
		curr = (curr<<1)%mod + int(s[i]-'0')
		// fmt.Println(curr)
		dict[curr] = true
	}
	// fmt.Println(dict)
	for i := 0; i < len(dict); i++ {
		if dict[i] == false {
			return false
		}
	}
	// fmt.Println(2<<(k-1), dict)
	return true
}

func hasAllCodes2602233(s string, k int) bool {
	if len(s) <= k {
		return false
	}
	dict := make([]bool, 2<<(k-1))
	curr := 0
	for i := 0; i < k; i++ {
		curr = curr<<1 | int(s[i]-'0')
	}
	dict[curr] = true
	count := 1
	mask := 1<<k - 1
	for i := k; i < len(s); i++ {
		curr = (curr<<1)&mask | int(s[i]-'0')
		if !dict[curr] {
			dict[curr] = true
			count += 1
		}
	}
	return count == 1<<k
}
