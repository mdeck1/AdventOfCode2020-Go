package level3

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
)

func Part1() {
	lines := utils.ReadFile("level3\\level3data.txt")
	var grid [][]bool = make([][]bool, 0)
	for i, line := range lines {
		grid = append(grid, make([]bool, 0))
		for j := 0; j < len(line); j++ {
			grid[i] = append(grid[i], line[j] == '#')
		}
	}
	fmt.Println(countTreesOnPath(3, 1, grid))
}

func Part2() {
	lines := utils.ReadFile("level3\\level3data.txt")
	var grid [][]bool = make([][]bool, 0)
	for i, line := range lines {
		grid = append(grid, make([]bool, 0))
		for j := 0; j < len(line); j++ {
			grid[i] = append(grid[i], line[j] == '#')
		}
	}
	val := countTreesOnPath(1, 1, grid)
	val *= countTreesOnPath(3, 1, grid)
	val *= countTreesOnPath(5, 1, grid)
	val *= countTreesOnPath(7, 1, grid)
	val *= countTreesOnPath(1, 2, grid)
	fmt.Println(val)
}

func countTreesOnPath(xDelta, yDelta int, grid [][]bool) int {
	x, y, count := 0, 0, 0
	for ; y < len(grid); y = y + yDelta {
		if grid[y][x] {
			count++
		}
		x = (x + xDelta) % len(grid[0])
	}
	return count
}
