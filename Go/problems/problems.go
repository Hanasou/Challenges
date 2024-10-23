package problems

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Helper function for returning max int
func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Helper function to find index of int array
func IndexOfArray(arr []int, ele int) int {
	for i, v := range arr {
		if v == ele {
			return i
		}
	}
	return -1
}

func Reverse(str string) string {
	new := ""

	for i := len(str) - 1; i >= 0; i-- {
		new += string(str[i])
	}

	return new
}

func ReverseRecursive(str string) string {
	return ReverseRecursiveHelper(str, "", len(str)-1)
}

func ReverseRecursiveHelper(str string, val string, index int) string {
	if index < 0 {
		return val
	}
	val += string(str[index])
	return ReverseRecursiveHelper(str, val, index-1)
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

func PalindromeRecursive(str string) bool {
	return PalindromeRecursiveHelper(str, 0, len(str)-1)
}

func PalindromeRecursiveHelper(str string, front int, back int) bool {
	if front > back {
		return true
	}
	if str[front] != str[back] {
		return false
	}
	return PalindromeRecursiveHelper(str, front+1, back-1)
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

// Return number of vowels in string
func Vowels(s string) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	vCount := 0

	for _, v := range strings.ToLower(s) {
		if _, ok := vowels[v]; ok {
			vCount++
		}
	}
	return vCount
}

// This function accepts an integer n and returns an nxn spiral matrix
// e.g Matrix(2) => [[1,2], [4,3]]
// Matrix(3) => [[1,2,3], [8,9,4], [7,6,5]]
func Matrix(n int) [][]int {
	// Initialize matrix
	solution := [][]int{}
	// Fill matrix with 0s
	for i := 1; i <= n; i++ {
		curr := []int{}
		for j := 1; j <= n; j++ {
			curr = append(curr, 0)
		}
		solution = append(solution, curr)
	}

	// Counter will track the number we place into the matrix
	counter := 1

	// Initialize several variables that represent the bounds of the matrix
	// We'll be changing the bounds while we iterate
	startRow, startCol := 0, 0
	endRow, endCol := n-1, n-1

	for startRow <= endRow && startCol <= endCol {
		// Fill top row
		for i := startCol; i <= endCol; i++ {
			solution[startRow][i] = counter
			counter++
		}
		startRow++
		// Fill right col
		for j := startRow; j <= endRow; j++ {
			solution[j][endCol] = counter
			counter++
		}
		endCol--
		// Fill bottom row
		for k := endCol; k >= startCol; k-- {
			solution[endRow][k] = counter
			counter++
		}
		endRow--
		// Fill left col
		for l := endRow; l >= startRow; l-- {
			solution[l][endCol] = counter
			counter++
		}
		startCol++
	}
	return solution
}

func FibIterative(n int) int {
	// The first number in the fibonacci series is 0
	first := 0
	// The second is 1
	second := 1
	total := first + second

	for i := 3; i <= n; i++ {
		total = first + second
		first = second
		second = total
	}
	return total
}

func FibRecursive(n int) int {
	mem := []int{}
	for i := 0; i <= n; i++ {
		mem = append(mem, -1)
	}
	return FibRecursiveHelper(n, mem)
}

func FibRecursiveHelper(n int, mem []int) int {
	if n < 2 {
		return n
	}
	if mem[n] == -1 {
		mem[n] = FibRecursiveHelper(n-1, mem) + FibRecursiveHelper(n-2, mem)
	}
	return mem[n]
}

func MaxProfit(prices []int) int {
	// Choose a single day to buy, and a single day after that to sell
	// Return the maximum profit you can achieve from this transaction
	maxProfit := 0
	minStock := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		currStock := prices[i]
		maxProfit = MaxInt(maxProfit, currStock-minStock)
		minStock = MinInt(currStock, minStock)
	}
	return maxProfit
}

// nums is rotated around an unknown pivot
// Return index of target, or -1 if the target is not in nums
func searchRotated(nums []int, target int) int {
	return 0
}
