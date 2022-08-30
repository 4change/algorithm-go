package problem0344

func reverseString(s []byte) {
    for left, right := 0, len(s)-1; left < right; left++ {
        s[left], s[right] = s[right], s[left]
        right--
    }
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}