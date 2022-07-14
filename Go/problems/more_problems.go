package problems

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
