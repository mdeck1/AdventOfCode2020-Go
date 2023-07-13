package utils

import (
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) []string {
	return ReadFileAndSplit(filepath, "\r\n")

}

func ReadFileAndSplit(filepath string, splitter string) []string {
	dat, err := os.ReadFile("C:\\Users\\malco\\ProgrammingProjects\\AdventOfCode2020-Go\\" + filepath)
	Check(err)
	return strings.Split(string(dat), splitter)
}

func StringSliceToIntSlice(string_list []string) []int {
	num_list := make([]int, 0, len(string_list))
	for _, str := range string_list {
		num, _ := strconv.Atoi(str)
		num_list = append(num_list, num)
	}
	return num_list
}
