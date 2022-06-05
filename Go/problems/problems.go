package problems

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
		if _, ok := counts[charString]; ok {
			counts[charString]++
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

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(n)
		}
	}
}

// Chunk takes in an array and a size and will split the input array into
// subarrays each with length size. Remainders will simply be added into another array
func Chunk(arr []int, size int) [][]int {
	result := [][]int{}
	current := []int{}
	for i, v := range arr {
		current = append(current, v)
		if (i+1)%size == 0 || i == len(arr)-1 {
			result = append(result, current)
			current = []int{}
		}
	}
	return result
}

func IsAnagram(a string, b string) bool {
	// Regex to remove all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-z0-9]+")
	if err != nil {
		fmt.Println(err)
	}

	// Remove all non-alphanumeric characters and make strings lowercase
	procA := strings.ToLower(reg.ReplaceAllString(a, ""))
	procB := strings.ToLower(reg.ReplaceAllString(b, ""))

	if len(procA) != len(procB) {
		return false
	}
	// Compare the two transformed strings
	aMap := map[byte]int{}
	bMap := map[byte]int{}
	c := 0

	for c < len(procA) && c < len(procB) {
		aRune := procA[c]
		bRune := procB[c]
		if _, ok := aMap[aRune]; ok {
			aMap[aRune]++
		} else {
			aMap[aRune] = 1
		}

		if _, ok := bMap[bRune]; ok {
			bMap[bRune]++
		} else {
			bMap[bRune] = 1
		}
		c++
	}

	return reflect.DeepEqual(aMap, bMap)
}

func Steps(n int) {
	for i := 1; i <= n; i++ {
		row := ""
		for j := 1; j <= i; j++ {
			row += "#"
		}
		for k := 1; k <= n-i; k++ {
			row += "@"
		}
		fmt.Println(row)
	}
}

func StepsRecursive(n int) {
	if n == 0 {
		return
	}

	stepsRecursiveHelper(n, n)
}

func stepsRecursiveHelper(curr int, total int) {
	if curr == 0 {
		return
	}
	row := ""
	for i := 1; i <= total-curr+1; i++ {
		row += "#"
	}
	for j := 1; j <= curr-1; j++ {
		row += "@"
	}
	fmt.Println(row)
	stepsRecursiveHelper(curr-1, total)
}

func Pyramid(n int) {
	for i := 1; i <= n; i++ {
		row := ""
		// First set of empty space
		for j := 1; j <= n-i; j++ {
			row += "@"
		}
		// First half of pound signs
		for k := 1; k <= i; k++ {
			row += "#"
		}
		// Second half of pound signs
		for l := 1; l <= i-1; l++ {
			row += "#"
		}
		// Second set of empty space
		for m := 1; m <= n-i; m++ {
			row += "@"
		}
		fmt.Println(row)
	}
}
