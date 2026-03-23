package _026

func removeStar(s string) string {
	var str []rune

	for i := 0; i < len(s); i++ {
		if s[i] == '*' && len(s) > 0 {
			str = str[:len(str)-1]
			continue
		}
		println(string(str))
		str = append(str, rune(s[i]))
	}

	return string(str)
}
