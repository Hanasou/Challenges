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

func MaxChar(str string) string {
	counts := map[string]int{}
	for _, v := range str {
		charString := string(v)
		if i, ok := counts[charString]; ok {
			counts[charString] = i + 1
		} else {
			counts[charString] = 1
		}
	}

	// iterate over map for largest value
	maxChar := ""
	maxCount := 0
	for k, v := range counts {
		if v > maxCount {
			maxCount = v
			maxChar = string(k)
		}	
	}
	return maxChar
}