package main

func GenerateParenthesis(n int) []string {
	out := []string{}
	backtrack("", n, n, &out)
	return out
}

func backtrack(current string, openN, closeN int, out *[]string) {
	if openN == 0 && closeN == 0 {
		*out = append(*out, current)
		return
	}
	if openN > 0 {
		backtrack(current+"(", openN-1, closeN, out)
	}
	if openN < closeN {
		backtrack(current+")", openN, closeN-1, out)
	}
}
