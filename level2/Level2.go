package level2

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"strconv"
	"strings"
)

func Part1() {
	lines := utils.ReadFile("level2\\level2data.txt")
	count := 0
	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")
		min, err := strconv.Atoi(strings.Split(words[0], "-")[0])
		utils.Check(err)
		max, err := strconv.Atoi(strings.Split(words[0], "-")[1])
		utils.Check(err)
		letter := words[1][:1]
		occurrences := strings.Count(words[2], letter)
		if occurrences <= max && occurrences >= min {
			count++
		}
	}
	fmt.Println(count)
}

func Part2() {
	lines := utils.ReadFile("level2\\level2data.txt")
	count := 0
	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")
		first, err := strconv.Atoi(strings.Split(words[0], "-")[0])
		utils.Check(err)
		first = first - 1
		second, err := strconv.Atoi(strings.Split(words[0], "-")[1])
		utils.Check(err)
		second = second - 1
		letter := words[1][0]
		if (words[2][first] == letter && words[2][second] != letter) || (words[2][first] != letter && words[2][second] == letter) {
			count++
		}
	}
	fmt.Println(count)
}
