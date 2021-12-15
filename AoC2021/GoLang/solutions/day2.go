package solutions

import (
	"AdventOfCode/AoC2021/GoLang/utils"
	"strconv"
	"strings"
)

func D2P1() (int, error) {
	text, err := utils.ReadFile("day2.txt")
	if err != nil {
		return 0, err
	}
	horizontal := 0
	depth := 0
	for _, v := range text {
		tokens := strings.Fields(v)
		rule := tokens[0]
		movement, _ := strconv.Atoi(tokens[1])
		switch rule {
		case "forward":
			horizontal += movement
		case "down":
			depth += movement
		case "backward":
			horizontal -= movement
		case "up":
			depth -= movement
		}
	}

	return horizontal * depth, nil
}

func D2P2() (int, error) {
	text, err := utils.ReadFile("day2.txt")
	if err != nil {
		return 0, err
	}
	horizontal := 0
	depth := 0
	aim := 0
	for _, v := range text {
		tokens := strings.Fields(v)
		rule := tokens[0]
		movement, _ := strconv.Atoi(tokens[1])
		switch rule {
		case "forward":
			horizontal += movement
			depth += aim * movement
		case "down":
			aim += movement
		case "backward":
			horizontal -= movement
		case "up":
			aim -= movement
		}
	}

	return horizontal * depth, nil
}
