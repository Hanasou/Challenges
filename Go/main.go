package main

import (
	"fmt"
	"sandbox/problems"
)

func main() {
	n := problems.ConvertToRanges([]int{0, 1, 3, 5, 50, 75}, 0, 100)
	fmt.Println(n)
}
