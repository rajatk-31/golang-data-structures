package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	seen := make([][]bool, len(grid))
	for i := range seen {
		seen[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !seen[i][j] {
				if c := CountIsland(grid, seen, i, j); c > max {
					max = c
				}
			}
		}
	}
	return max
}

func CountIsland(grid [][]int, seen [][]bool, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}
	if seen[i][j] {
		return 0
	} else {
		sum := 1
		seen[i][j] = true
		if i+1 < len(grid) {
			sum = sum + CountIsland(grid, seen, i+1, j)
		}
		if i-1 >= 0 {
			sum = sum + CountIsland(grid, seen, i-1, j)
		}
		if j+1 < len(grid[0]) {
			sum = sum + CountIsland(grid, seen, i, j+1)
		}
		if j-1 >= 0 {
			sum = sum + CountIsland(grid, seen, i, j-1)
		}
		return sum
	}
}
