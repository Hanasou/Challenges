package main

import (
	"reflect"
	"strconv"
	"strings"
)

func anagram(s1, s2 string) bool {
	s1, s2 = strings.ToLower(s1), strings.ToLower(s2)
	s1, s2 = strings.Replace(s1, " ", "", -1), strings.Replace(s2, " ", "", -1)
	var s1counts = map[rune]int{}
	var s2counts = map[rune]int{}

	for _, c := range s1 {
		if _, ok := s1counts[c]; ok {
			s1counts[c]++
		} else {
			s1counts[c] = 1
		}
	}
	for _, c := range s2 {
		if _, ok := s2counts[c]; ok {
			s2counts[c]++
		} else {
			s2counts[c] = 1
		}
	}
	return reflect.DeepEqual(s1counts, s2counts)
}

// Return all pairs that sum up to k
func pairSum(arr []int, k int) [][2]int {
	pairs := [][2]int{}
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				pairs = append(pairs, [2]int{arr[i], arr[j]})
			}
		}
	}
	return pairs
}

// Returns the missing number
func missingNumber(arr1, arr2 []int) int {
	var arrCounts = map[int]int{}
	for _, v := range arr1 {
		if _, ok := arrCounts[v]; ok {
			arrCounts[v]++
		} else {
			arrCounts[v] = 1
		}
	}
	for _, v := range arr2 {
		arrCounts[v]--
	}

	for k, v := range arrCounts {
		if v != 0 {
			return k
		}
	}
	return -1
}

func reverseSentence(s string) string {
	s = strings.TrimSpace(s)
	sSlice := strings.Split(s, " ")
	i := 0
	j := len(sSlice) - 1
	for i < j {
		sSlice[i], sSlice[j] = sSlice[j], sSlice[i]
		i++
		j--
	}
	retString := strings.Join(sSlice, " ")
	return retString
}

func stringCompression(s string) string {
	var counts = map[rune]int{}
	compressed := ""
	for _, c := range s {
		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c] = 1
		}
	}
	for k, v := range counts {
		compressed += string(k) + strconv.Itoa(v)
	}
	return compressed
}

func allUnique(s string) bool {
	var counts = map[rune]bool{}
	for _, c := range s {
		if _, ok := counts[c]; ok {
			return false
		}
		counts[c] = true
	}
	return true
}
