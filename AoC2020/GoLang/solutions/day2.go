package solutions

import (
	"AdventOfCode/AoC2020/GoLang/utils"
	"strconv"
	"strings"
)

func D2P1() (int, error) {
	text, err := utils.ReadFile("day2.txt")
	if err != nil {
		return 0, err
	}
	total := 0
	for _, v := range text {
		tokens := strings.Fields(v)
		rules := strings.Split(tokens[0], "-")
		min, _ := strconv.Atoi(rules[0])
		max, _ := strconv.Atoi(rules[1])
		letter := tokens[1][:1]
		count := strings.Count(tokens[2], letter)
		if count <= max && count >= min {
			total++
		}
	}

	return total, nil
}

func D2P2() (int, error) {
	text, err := utils.ReadFile("day2.txt")
	if err != nil {
		return 0, err
	}
	total := 0
	for _, v := range text {
		tokens := strings.Fields(v)
		rules := strings.Split(tokens[0], "-")
		min, _ := strconv.Atoi(rules[0])
		max, _ := strconv.Atoi(rules[1])
		letter := tokens[1][:1]
		if bool(string(tokens[2][min-1]) == letter) != bool(string(tokens[2][max-1]) == letter) {
			total++
		}
	}

	return total, nil
}
