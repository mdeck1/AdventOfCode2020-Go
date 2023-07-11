package level8

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"strconv"
	"strings"
)

func Part1() {
	lines := utils.ReadFile("level8\\level8data.txt")
	acc, _ := accValueAtRecurseOrExit(lines)
	fmt.Println(acc)
}

func Part2() {
	lines := utils.ReadFile("level8\\level8data.txt")
	for i, _ := range lines {
		words := strings.Fields(lines[i])
		if words[0] == "nop" {
			lines[i] = "jmp " + words[1]
			acc, recursed := accValueAtRecurseOrExit(lines)
			if !recursed {
				fmt.Println(acc)
				return
			}
			lines[i] = "nop " + words[1]
		} else if words[0] == "jmp" {
			lines[i] = "nop " + words[1]
			acc, recursed := accValueAtRecurseOrExit(lines)
			if !recursed {
				fmt.Println(acc)
				return
			}
			lines[i] = "jmp " + words[1]
		}
	}
}

func accValueAtRecurseOrExit(lines []string) (int, bool) {
	acc := 0
	seen := make([]bool, len(lines))
	currLine := 0
	for currLine < len(lines) {
		if seen[currLine] {
			break
		}
		seen[currLine] = true
		words := strings.Fields(lines[currLine])
		if words[0] == "acc" {
			if words[1][0] == '+' {
				add, _ := strconv.Atoi(words[1][1:])
				acc += add
			} else {
				sub, _ := strconv.Atoi(words[1][1:])
				acc -= sub
			}
			currLine++
		} else if words[0] == "jmp" {
			if words[1][0] == '+' {
				add, _ := strconv.Atoi(words[1][1:])
				currLine += add
			} else {
				sub, _ := strconv.Atoi(words[1][1:])
				currLine -= sub
			}
		} else {
			currLine++
		}
	}
	return acc, currLine < len(lines)
}
