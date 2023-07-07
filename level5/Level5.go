package level5

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"sort"
)

func Part1() {
	lines := utils.ReadFile("level5\\level5data.txt")
	max := 0
	for _, line := range lines {
		val := stringToSeatId(line)
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}

func Part2() {
	lines := utils.ReadFile("level5\\level5data.txt")
	var seats = make([]int, 0)
	for _, line := range lines {
		val := stringToSeatId(line)
		seats = append(seats, val)
	}
	sort.Ints(seats)
	for i := 0; i < len(seats); i++ {
		if seats[i] == seats[i+1]-2 {
			fmt.Println(seats[i] + 1)
			break
		}
	}
}

func stringToSeatId(line string) int {
	row := 0
	for i := 0; i < 7; i++ {
		row = row << 1
		if line[i] == 'B' {
			row += 1
		}
	}
	col := 0
	for i := 7; i < 10; i++ {
		col = col << 1
		if line[i] == 'R' {
			col += 1
		}
	}
	return row*8 + col
}
