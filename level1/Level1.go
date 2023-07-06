package level1

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"strconv"
)

func Part1() {
	lines := utils.ReadFile("level1\\level1data.txt")
	var numbers []int
	for i := 0; i < len(lines); i++ {
		num, err := strconv.Atoi(lines[i])
		utils.Check(err)
		numbers = append(numbers, num)
	}

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if numbers[i]+numbers[j] == 2020 {
				fmt.Println(numbers[i] * numbers[j])
			}
		}
	}
}

func Part2() {
	lines := utils.ReadFile("level1\\level1data.txt")
	var numbers []int
	for i := 0; i < len(lines); i++ {
		num, err := strconv.Atoi(lines[i])
		utils.Check(err)
		numbers = append(numbers, num)
	}

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			for k := j + 1; k < len(lines); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i] * numbers[j] * numbers[k])
				}
			}
		}
	}
}
