pakage main

func smallestPalindrome260728(s string) string {
    str := []byte(s)
    sort.Slice(str[:len(str)/2], func(i, j int) bool { return str[i]<=str[j]})
    for l, r := 0, len(str)-1; l < r; l, r = l+1, r-1 {
        str[r] = str[l]
    }
    return string(str)
}