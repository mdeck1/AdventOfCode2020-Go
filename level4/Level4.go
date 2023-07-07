package level4

import (
	"encoding/hex"
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"strconv"
	"strings"
)

func Part1() {
	lines := utils.ReadFileAndSplit("level4\\level4data.txt", "\r\n\r\n")
	count := 0
	for _, line := range lines {
		if containsAllFields(line) {
			count++
		}
	}
	fmt.Println(count)
}

func Part2() {
	lines := utils.ReadFileAndSplit("level4\\level4data.txt", "\r\n\r\n")
	count := 0
	for _, line := range lines {
		if !containsAllFields(line) {
			continue
		}
		words := strings.Fields(line)
		valid := true
		for _, word := range words {
			if strings.Contains(word, "byr:") {
				val, err := strconv.Atoi(word[4:])
				utils.Check(err)
				if val < 1920 || val > 2002 {
					valid = false
					printIssue(line, "byr")
					break
				}
			}
			if strings.Contains(word, "iyr:") {
				val, err := strconv.Atoi(word[4:])
				utils.Check(err)
				if val < 2010 || val > 2020 {
					valid = false
					printIssue(line, "iyr")
					break
				}
			}
			if strings.Contains(word, "eyr:") {
				val, err := strconv.Atoi(word[4:])
				utils.Check(err)
				if val < 2020 || val > 2030 {
					valid = false
					printIssue(line, "eyr")
					break
				}
			}
			if strings.Contains(word, "hgt:") {
				if !strings.Contains(word, "cm") && !strings.Contains(word, "in") {
					valid = false
					printIssue(line, "hgt")
					break
				}
				val, err := strconv.Atoi(word[4 : len(word)-2])
				utils.Check(err)
				if strings.Contains(word, "cm") {
					if val < 150 || val > 193 {
						valid = false
						printIssue(line, "hgt")
						break
					}
				} else if strings.Contains(word, "in") {
					if val < 59 || val > 76 {
						valid = false
						printIssue(line, "hgt")
						break
					}
				}
			}
			if strings.Contains(word, "hcl:") {
				if len(word) != 11 || word[4] != '#' {
					valid = false
					printIssue(line, "hcl")
					break
				}
				_, err := hex.Decode(make([]byte, hex.DecodedLen(len(word[5:]))), []byte(word[5:]))
				if err != nil || len(word[5:]) != 6 {
					valid = false
					printIssue(line, "hcl")
					break
				}
			}
			if strings.Contains(word, "ecl:") {
				if word[4:] != "amb" && word[4:] != "blu" && word[4:] != "brn" && word[4:] != "gry" && word[4:] != "grn" && word[4:] != "hzl" && word[4:] != "oth" {
					valid = false
					printIssue(line, "ecl")
					break
				}
			}
			if strings.Contains(word, "pid:") {
				_, err := strconv.Atoi(word[4:])
				if len(word) != 13 || err != nil {
					valid = false
					printIssue(line, "pid")
					break
				}
			}
		}
		if valid {
			count++
		}
	}
	fmt.Println(count)
}

func containsAllFields(line string) bool {
	if !strings.Contains(line, "byr:") {
		return false
	}
	if !strings.Contains(line, "iyr:") {
		return false
	}
	if !strings.Contains(line, "eyr:") {
		return false
	}
	if !strings.Contains(line, "hgt:") {
		return false
	}
	if !strings.Contains(line, "hcl:") {
		return false
	}
	if !strings.Contains(line, "ecl:") {
		return false
	}
	if !strings.Contains(line, "pid:") {
		return false
	}
	return true
}

func printIssue(line, field string) {
	if false {
		fmt.Printf("Issue with field %v in line: %v\n", field, line)
	}
}
