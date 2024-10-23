package problems

import (
	"fmt"
	"strconv"
)

type pair struct {
	x int
	y int
}

func numIslands(grid [][]int) int {
	// Make a grid of visited squares in our grid
	// We will use this to see where we still need to search
	visited := [][]bool{}

	// Fill the matrix with false values
	for i := 0; i < len(grid); i++ {
		row := []bool{}
		for j := 0; j < len(grid[0]); j++ {
			row = append(row, false)
		}
		visited = append(visited, row)
	}

	// Perform Breadth First Search when we hit a "1" on the grid
	// and an Undiscovered (False) location on corresponding visited matrix
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			p := pair{
				x: i,
				y: j,
			}
			if visited[i][j] == false && grid[i][j] == 1 {
				visited = bfs(grid, visited, p)
				islands++
			}
		}
	}

	return islands
}

func bfs(grid [][]int, visited [][]bool, point pair) [][]bool {
	q := []pair{}
	q = append(q, point)

	for len(q) > 0 {
		// Pop element in queue
		var p pair
		q, p = pop(q)
		// Extract x and y elements from pair
		x := p.x
		y := p.y
		// Mark visited
		visited[x][y] = true

		// Put adjacent pairs into the queue only if it's land and is in bound
		n := pair{
			x: x,
			y: y - 1,
		}
		e := pair{
			x: x + 1,
			y: y,
		}
		w := pair{
			x: x - 1,
			y: y,
		}
		s := pair{
			x: x,
			y: y + 1,
		}
		if isInBounds(visited, n) && grid[n.x][n.y] == 1 && visited[n.x][n.y] == false {
			q = append(q, n)
		}
		if isInBounds(visited, e) && grid[e.x][e.y] == 1 && visited[e.x][e.y] == false {
			q = append(q, e)
		}
		if isInBounds(visited, w) && grid[w.x][w.y] == 1 && visited[w.x][w.y] == false {
			q = append(q, w)
		}
		if isInBounds(visited, s) && grid[s.x][s.y] == 1 && visited[s.x][s.y] == false {
			q = append(q, s)
		}
	}
	return visited
}

func pop(q []pair) ([]pair, pair) {
	return q[:len(q)-1], q[len(q)-1]
}

func isInBounds(grid [][]bool, point pair) bool {
	if point.x < 0 || point.x >= len(grid) ||
		point.y < 0 || point.y >= len(grid[0]) {
		return false
	}
	return true
}

// Return all unique combinations of candidates where the chosen numbers sum up to target
// The same number may be chosen from candidates an unlimited number of times
// Two combinations are unique if the frequency of at least one of the chosen numbers
// is different.
func combinationSum(candidates []int, target int) [][]int {
	return nil
}

// Given an integer array nums, and integer k
// Return the k most frequent elements
// Example
// Input: nums = [1,1,1,2,2,3], k = 2
// Ouptut: [1,2]
// The k (2) most frequently occurring numbers in "nums" is 1 with 3 and 2, with 2.
// The algorithm must be better than O(nlogn), n being the size of nums
func topKFrequent(nums []int, k int) []int {
	// Try storing all the counts in a Map
	// Key = number, value is the frequency
	// How do we take the top k from the map?
	// First Idea: Find the max value of the map k times, excluding the previous max
	return nil
}

func uniquePaths(m int, n int) int {
	// Each position on this grid has a number of unique paths to get to that spot
	// The robot can only move down, or right, so the number of ways to get to one particular spot is:
	// (# of ways it takes to get to the north spot) + (# of ways it takes to get to the west spot)

	// Initialize m x n grid
	rows := [][]int{}
	for i := 0; i < m; i++ {
		columns := []int{}
		for j := 0; j < n; j++ {
			columns = append(columns, 0)
		}
		rows = append(rows, columns)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					rows[i][j] = 1
				} else {
					rows[i][j] = rows[i][j-1]
				}
			} else if j == 0 {
				rows[i][j] = rows[i-1][j]
			} else {
				rows[i][j] = rows[i-1][j] + rows[i][j-1]
			}
		}
	}
	return rows[m][n]
}

func uniquePathsRecurse(grid [][]int, m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	if grid[m][n] == 0 {
		grid[m][n] = uniquePathsRecurse(grid, m-1, n) + uniquePathsRecurse(grid, m, n-1)
	}
	return grid[m][n]
}

// You are given an array of non overlapping intervals
// intervals[i] = [start, end]
// Intervals is sorted by start
// You are given a new interval [start, end]
// Insert newInterval into intervals such that:
// intervals is still sorted in ascending order by start
// and intervals does not have any overlapping intervals (so you want to merge the intervals)
func insertInterval(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	for i, currInterval := range intervals {
		if newInterval[0] > currInterval[1] {
			// [1,3],[4,5],[9,10], <- [6,8]
			result = append(result, currInterval)
		} else if newInterval[0] > currInterval[1] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else {
			newInterval = []int{MinInt(currInterval[0], newInterval[0]), MaxInt(currInterval[1], newInterval[1])}
		}
	}

	result = append(result, newInterval)
	return result
}

// You are given a list of numbers.
// Each number represents the amount of money inside a house
// You cannot rob two adjacent houses
// Return the max amount of money you can rob
// [1,2,3,1]
// Answer: 4
// Rob house 1, then rob house 3. You can't rob house 2 and 3 because they're adjacent
func rob(nums []int) int {
	// Maintain a list of the maximum amount of money you can rob on each night
	mem := []int{}
	for i, v := range nums {
		if i == 0 {
			mem = append(mem, v)
		} else if i == 1 {
			mem = append(mem, MaxInt(mem[0], v))
		} else {
			mem = append(mem, MaxInt(mem[i-1], mem[i-2]))
		}
	}
	return mem[len(nums)-1]
}

// There are a total of numCourses you have to take, labeled from 0 to numCourses - 1
// Given an array of prerequisites, where prerequisites[i] = [a, b] indicates you must take course b first if you want to take course a
// For example, pair [0,1] indicates that in order to take course 0, you have to take course 1
// Return true if you can finish all courses, otherwise return false
func canFinish(numCourses int, prerequisites [][]int) bool {
	q := []int{}
	visited := map[int]bool{}
	first := prerequisites[0][0]
	q = append(q, first)

	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]                // represents popping the queue
		if _, ok := visited[curr]; ok { // Cycle detected
			return false
		}
		visited[curr] = true
		for _, v := range prerequisites {
			if v[0] == curr {
				q = append(q, v[1])
			}
		}
	}
	return len(visited) == numCourses
}

func canFinishDfs(numCourses int, prerequisites [][]int, visited map[int]bool, curr int) bool {
	if numCourses == len(visited) {
		return true
	}
	if _, ok := visited[curr]; ok {
		return false
	}
	visited[curr] = true
	for _, v := range prerequisites {
		if v[0] == curr {
			return canFinishDfs(numCourses, prerequisites, visited, v[1])
		}
	}
	return true
}

// Given a sorted input array of integers, and a lower and upper bounds
// Return an array of strings that maps to each interval in the array
// Also include the lower and upper bounds in there as well
// E.g. [0,1,3,5,50,75], lower = 0, upper = 100
// return ["2", "4", "6->49", "51->74", "76->99"]
func ConvertToRanges(nums []int, lower int, upper int) []string {
	// First thing let's try doing is insert lower and upper into the nums array
	// Then we look at the space in between each element in the nums array
	// And we can create a string that represents the interval between two numbers

	lowerInserted := false
	upperInserted := false

	solution := []string{}

	for i, v := range nums {
		// Let's first pass through every element that's lower than our lower bound
		// and greater than our upper bound
		if v <= lower || v >= upper {
			continue
		}
		// If we found a value that is greater than our lower bound, we can insert it
		if v > lower && !lowerInserted {
			// prepend lower
			nums = prependItem(nums, i, v)
			i += 1
			lowerInserted = true
		}
		if v > upper && !upperInserted {
			// prepend upper
			nums = prependItem(nums, i, v)
			i += 1
			upperInserted = true
		}

		if i < len(nums) {
			// Build interval
			x := ""
			lowerV := nums[i] + 1
			upperV := (nums[i+1]) - 1

			fmt.Println("Curr i: ", i)
			fmt.Println("Lowerv: ", lowerV)
			fmt.Println("Upperv: ", upperV)

			// If they are the same number
			if lowerV > upperV {
				continue
			} else if lowerV == upperV {
				x += strconv.Itoa(lowerV)
			} else {
				x += strconv.Itoa(lowerV) + "->" + strconv.Itoa(upperV)
			}
			solution = append(solution, x)
		} else {
			continue
		}
	}
	return solution
}

func prependItem(nums []int, index int, value int) []int {
	nums = append(nums[:index+1], nums[index:]...)
	nums[index] = value
	return nums
}
