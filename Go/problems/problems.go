package problems

import (
	"fmt"
	"strconv"
	"math"
)

func Reverse(str string) string {
	new := ""

	for i := len(str) - 1; i >= 0; i-- {
		new += string(str[i])
	}

	return new
}

func Palindrome(str string) bool {
	front := 0
	back := len(str) - 1
	for front <= back {
		if str[front] != str[back] {
			return false
		}
		front++
		back--
	}
	return true
}

func ReverseInt(n int) int {
	isNegative := n < 0
	nString := strconv.Itoa(int(math.Abs(float64(n))))
	new := ""
	if isNegative {
		new += "-"
	}

	for i := len(nString) - 1; i >= 0; i-- {
		new += string(nString[i])
	}

	newInt, err := strconv.Atoi(new)
	if err != nil {
		fmt.Println(err)
	}

	return newInt
}