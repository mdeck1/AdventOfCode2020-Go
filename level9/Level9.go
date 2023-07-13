package level9

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"strconv"
)

func Part1() {
	lines := utils.ReadFile("level9\\level9data.txt")
	const PREAMBLE_LENGTH int = 25
	num_list := make([]int, 0, 25)
	num_map := make(map[int]int)
	for i := 0; i < PREAMBLE_LENGTH; i++ {
		num, _ := strconv.Atoi(lines[i])
		num_list = append(num_list, num)
		count := num_map[num]
		num_map[num] = count + 1
	}
	for i := PREAMBLE_LENGTH; i < len(lines); i++ {
		target_val, _ := strconv.Atoi(lines[i])
		if !checkForTargetVal(target_val, num_list, num_map) {
			fmt.Println("First misprint: " + fmt.Sprint(target_val))
			return
		}
		num_to_remove := num_list[0]
		num_map[num_to_remove] = num_map[num_to_remove] - 1
		num_map[target_val] = num_map[target_val] + 1
		num_list = num_list[1:]
		num_list = append(num_list, target_val)
	}
}

func Part2() {
	lines := utils.ReadFile("level9\\level9data.txt")
	magic_num := 104054607
	num_list := make([]int, 0, 1000)
	for i := 0; i < len(lines); i++ {
		num, _ := strconv.Atoi(lines[i])
		num_list = append(num_list, num)
	}
	a, b := findRangeSummingToTarget(magic_num, num_list)
	fmt.Println("Encryption weakness: " +
		fmt.Sprint(sumOfSmallestAndLargestWithinBounds(num_list, a, b)))
}

func checkForTargetVal(target_val int, num_list []int, num_map map[int]int) bool {
	for _, currNum := range num_list {
		if ((target_val - currNum) == currNum) && num_map[target_val-currNum] > 1 {
			return true
		} else if num_map[target_val-currNum] > 0 {
			return true
		}
	}
	return false
}

// This is O(n^2), could be better done (O(n)) using a sliding-window
// There are only 1000 entries, though, so who cares :P
func findRangeSummingToTarget(target_val int, num_list []int) (int, int) {
	for i := 0; i < len(num_list); i++ {
		curr_sum := num_list[i]
		j := i + 1
		for ; j < len(num_list) && curr_sum < target_val; j++ {
			curr_sum += num_list[j]
		}
		if curr_sum == target_val {
			return i, j - 1
		}
	}
	// Should never happen
	return -1, -1
}

func sumOfSmallestAndLargestWithinBounds(num_list []int, min int, max int) int {
	min_val := num_list[min]
	max_val := num_list[min]
	for i := min; i <= max; i++ {
		if num_list[i] < min_val {
			min_val = num_list[i]
		}
		if num_list[i] > max_val {
			max_val = num_list[i]
		}
	}
	return min_val + max_val
}
