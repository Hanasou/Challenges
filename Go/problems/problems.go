package problems

func Reverse(str string) string {
	new := ""

	for i := len(str) - 1; i >= 0; i-- {
		new += string(str[i])
	}

	return new
}
