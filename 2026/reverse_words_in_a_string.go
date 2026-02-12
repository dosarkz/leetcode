package _026

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	words := []byte(s)
	var result []byte
	var word []byte

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == ' ' {
			if len(word) > 0 {
				if len(result) > 0 {
					result = append(result, ' ')
				}

				result = append(result, word...)
				word = []byte{}
			}
			continue
		}

		word = append([]byte{words[i]}, word...)

		if i == 0 {
			if len(result) > 0 {
				result = append(result, ' ')
			}
			result = append(result, word...)
		}

	}

	return string(result)
}
