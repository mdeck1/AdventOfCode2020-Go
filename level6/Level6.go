package level6

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
)

func Part1() {
	lines := utils.ReadFile("level6\\level6data.txt")
	count := 0
	for i := 0; i < len(lines); i++ {
		set := map[rune]struct{}{}
		for i < len(lines) && lines[i] != "" {
			for _, c := range lines[i] {
				set[c] = struct{}{}
			}
			i++
		}
		count += len(set)
	}
	fmt.Println(count)
}

func Part2() {
	lines := utils.ReadFile("level6\\level6data.txt")
	count := 0
	for i := 0; i < len(lines); i++ {
		set := make(map[rune]int)
		howManyPeople := 0
		for i < len(lines) && lines[i] != "" {
			howManyPeople++
			for _, c := range lines[i] {
				curr := set[c]
				set[c] = curr + 1
			}
			i++
		}
		for _, value := range set {
			if value == howManyPeople {
				count++
			}
		}
	}
	fmt.Println(count)
}
