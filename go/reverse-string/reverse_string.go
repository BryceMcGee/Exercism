package reverse

func Reverse(input string) string {
	if len(input) == 0 {
		return ""
	}

	r := []rune(input)
	reverse := ""

	for i := len(r) - 1; i >= 0; i-- {
		reverse += string(r[i])
	}
	return reverse
}