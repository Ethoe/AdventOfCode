package solutions

import (
	"AdventOfCode/AoC2020/GoLang/utils"
)

type point struct {
	x int
	y int
}

func add(p1, p2 point, size int) point {
	return point{(p1.x + p2.x) % size, p1.y + p2.y}
}

func simulator(text []string, movement point) (trees int) {
	position := point{0, 0}
	trees = 0

	for position.y < len(text) {
		if string(text[position.y][position.x]) == "#" {
			trees++
		}
		position = add(position, movement, len(text[position.x]))
	}
	return trees
}

func D3P1() (trees int, err error) {
	text, err := utils.ReadFile("day3.txt")
	if err != nil {
		return 0, err
	}

	return simulator(text, point{3, 1}), nil
}

func D3P2() (total int, err error) {
	text, err := utils.ReadFile("day3.txt")
	if err != nil {
		return 0, err
	}
	attempts := []point{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	total = 1
	for _, attempt := range attempts {
		total *= simulator(text, attempt)
	}
	return total, nil
}
