package utils

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) []string {
	dat, err := os.ReadFile("C:\\Users\\malco\\ProgrammingProjects\\AdventOfCode2020-Go\\" + filepath)
	Check(err)
	return strings.Split(string(dat), "\r\n")

}