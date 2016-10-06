package baby

// NameGenerator returns generator of baby names
func NameGenerator(first, second string) func() string {
	i, j := 0, 0
	f := []rune(first)
	s := []rune(second)
	return func() string {
		// 실제 array의 index 범위 내에 있을때만 가능
		// 종료 조건
		if i > len(f)-1 {
			return ""
		}
		value := string(f[i]) + string(s[j])
		// index add
		if j < len(s)-1 {
			j++
		} else {
			i++
			j = 0
		}
		return value
	}
}

// CallBack calls callback functions generate
func CallBack(first, second string, cb func(string)) {
	for _, f := range first {
		for _, s := range second {
			cb(string(f) + string(s))
		}
	}
}
