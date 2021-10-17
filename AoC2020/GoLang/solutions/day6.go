package solutions

import (
	"AdventOfCode/AoC2020/GoLang/utils"
)

func D6P1() (int, error) {
	text, err := utils.ReadFile("day6.txt")
	if err != nil {
		return 0, err
	}

	var set map[string]bool
	index := 0
	total := 0

	for index < len(text) {
		set = make(map[string]bool, 0)
		for true {
			if index >= len(text) || text[index] == "" {
				index++
				break
			} else {
				for _, letter := range text[index] {
					set[string(letter)] = true
				}
			}
			index++
		}
		total += len(set)
	}

	return total, nil
}

func D6P2() (int, error) {
	text, err := utils.ReadFile("day6.txt")
	if err != nil {
		return 0, err
	}

	var set map[string]bool
	var init bool
	index := 0
	total := 0

	for index < len(text) {
		set = make(map[string]bool, 0)
		init = false
		for true {
			if index >= len(text) || text[index] == "" {
				index++
				break
			} else {
				if !init {
					for _, letter := range text[index] {
						set[string(letter)] = true
					}
					init = true
				} else {
					newSet := make(map[string]bool, 0)
					for _, letter := range text[index] {
						newSet[string(letter)] = true
					}
					combine := intersection(set, newSet)
					set = combine
				}
			}
			index++
		}
		total += len(set)
	}

	return total, nil
}

func intersection(s1, s2 map[string]bool) (inter map[string]bool) {
	hash := make(map[string]bool)
	inter = make(map[string]bool)
	for e, _ := range s1 {
		hash[e] = true
	}
	for e, _ := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter[e] = true
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

//Remove dups from slice.
func removeDups(elements map[string]bool) (nodups map[string]bool) {
	encountered := make(map[string]bool)
	nodups = make(map[string]bool)
	for element, _ := range elements {
		if !encountered[element] {
			nodups[element] = true
			encountered[element] = true
		}
	}
	return
}
