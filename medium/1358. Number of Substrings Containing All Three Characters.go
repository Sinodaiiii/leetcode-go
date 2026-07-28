package medium

func numberOfSubstrings260630(s string) int {
    left := make([]int, 3)
    n := len(s)
    ans, curr := 0, 0
    l, r := 0, 0
    preL := -1
    for l <= r && r != n {
        for curr < 3 && r < n {
            rIndex := int(s[r]-'a')
            r += 1
            left[rIndex] += 1
            if left[rIndex] == 1 {
                curr += 1
            }
        }
        // fmt.Println(r, curr)
        for curr == 3 {
            lIndex := int(s[l]-'a')
            l += 1
            left[lIndex] -= 1
            if left[lIndex] == 0 {
                curr -= 1
                ans += (l-1-preL) * (n-r+1)
                preL = l-1
                // fmt.Println(l, r, preL, ans)
            }
        }
    }
    return ans
}