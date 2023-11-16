package graphs

func numIslands(ic IslandCount) int {
	grid := ic.Grid
	visited := make(map[[2]int]bool)
	var count int
	for i := range grid {
		for j := range grid[0] {
			if visited[[2]int{i, j}] {
				continue
			}
			if grid[i][j] == "1" {
				visitIsland(grid, i, j, visited)
				count++
			}
		}
	}

	return count
}

func visitIsland(grid [][]string, i int, j int, visited map[[2]int]bool) {
	if i < 0 || i >= len(grid) {
		return
	}
	if j < 0 || j >= len(grid[0]) {
		return
	}
	if visited[[2]int{i, j}] {
		return
	}
	if grid[i][j] == "0" {
		return
	}
	visited[[2]int{i, j}] = true
	visitIsland(grid, i+1, j, visited)
	visitIsland(grid, i-1, j, visited)
	visitIsland(grid, i, j+1, visited)
	visitIsland(grid, i, j-1, visited)
}
