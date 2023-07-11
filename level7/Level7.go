package level7

import (
	"fmt"
	"malcolmdeck/adventofcode2020/utils"
	"strconv"
	"strings"
)

type multibag struct {
	count    int
	bag_name string
}

func Part1() {
	lines := utils.ReadFile("level7\\level7data.txt")
	bag_map := make(map[string][]multibag)
	reverse_bag_map := make(map[string][]string)
	for _, line := range lines {
		words := strings.Fields(line)
		outer_bag_name := words[0] + " " + words[1]
		if len(words) == 7 {
			bag_map[outer_bag_name] = nil //[]multibag{}
			continue
		}
		bags := make([]multibag, len(words)/4-1)
		for i := 4; i < len(words); i += 4 {
			count, _ := strconv.Atoi(words[i])
			inner_bag_name := words[i+1] + " " + words[i+2]
			bags = append(bags, multibag{count: count, bag_name: inner_bag_name})
			reverse_bag_map[inner_bag_name] = append(reverse_bag_map[inner_bag_name], outer_bag_name)
		}
		bag_map[outer_bag_name] = bags
	}
	bags_to_check := make([]string, 0)
	bags_to_check = append(bags_to_check, reverse_bag_map["shiny gold"]...)
	seen_bags := map[string]struct{}{}
	count := 0
	for len(bags_to_check) != 0 {
		curr_bag := bags_to_check[0]
		bags_to_check = bags_to_check[1:]
		_, ok := seen_bags[curr_bag]
		if !ok {
			seen_bags[curr_bag] = struct{}{}
			count++
			// fmt.Println(curr_bag + " bags may contain shiny gold bags")
		}
		bags_to_check = append(bags_to_check, reverse_bag_map[curr_bag]...)
	}
	fmt.Println(count)
}

func Part2() {
	lines := utils.ReadFile("level7\\level7data.txt")
	bag_map := make(map[string][]multibag)
	reverse_bag_map := make(map[string][]string)
	for _, line := range lines {
		words := strings.Fields(line)
		outer_bag_name := words[0] + " " + words[1]
		if len(words) == 7 {
			bag_map[outer_bag_name] = nil //[]multibag{}
			continue
		}
		bags := make([]multibag, len(words)/4-1)
		for i := 4; i < len(words); i += 4 {
			count, _ := strconv.Atoi(words[i])
			inner_bag_name := words[i+1] + " " + words[i+2]
			bags = append(bags, multibag{count: count, bag_name: inner_bag_name})
			reverse_bag_map[inner_bag_name] = append(reverse_bag_map[inner_bag_name], outer_bag_name)
		}
		bag_map[outer_bag_name] = bags
	}
	fmt.Println(howManyBagsInThisBag("shiny gold", bag_map) - 1)
}

func howManyBagsInThisBag(name string, bag_map map[string][]multibag) int {
	inner_bag_count := 1
	for _, this_multibag := range bag_map[name] {
		inner_bag_count += this_multibag.count * howManyBagsInThisBag(this_multibag.bag_name, bag_map)
	}
	return inner_bag_count
}
