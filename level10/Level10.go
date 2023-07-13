package level10

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"sort"
)

func Part1() {
	lines := utils.ReadFile("level10\\level10data.txt")
	num_list := utils.StringSliceToIntSlice(lines)
	num_list = append(num_list, 0)
	sort.Ints(num_list)
	num_list = append(num_list, num_list[len(num_list)-1]+3)
	diff_count := []int{0, 0, 0}
	for i := 1; i < len(num_list); i++ {
		diff_count[num_list[i]-num_list[i-1]-1]++
	}
	fmt.Println(diff_count[0] * diff_count[2])
}

func Part2() {
	lines := utils.ReadFile("level10\\level10data.txt")
	num_list := utils.StringSliceToIntSlice(lines)
	num_list = append(num_list, 0)
	sort.Ints(num_list)
	num_list = append(num_list, num_list[len(num_list)-1]+3)

	paths_to_end_count := make([]int, len(num_list), len(num_list))
	paths_to_end_count[len(paths_to_end_count)-1] = 1
	for i := len(paths_to_end_count) - 2; i >= 0; i-- {
		if i+1 < len(paths_to_end_count) && num_list[i+1]-num_list[i] < 4 {
			paths_to_end_count[i] += paths_to_end_count[i+1]
		}
		if i+2 < len(paths_to_end_count) && num_list[i+2]-num_list[i] < 4 {
			paths_to_end_count[i] += paths_to_end_count[i+2]
		}
		if i+3 < len(paths_to_end_count) && num_list[i+3]-num_list[i] < 4 {
			paths_to_end_count[i] += paths_to_end_count[i+3]
		}
	}
	fmt.Println(paths_to_end_count[0])
}
