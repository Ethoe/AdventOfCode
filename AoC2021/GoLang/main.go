package main

import (
	"AdventOfCode/AoC2021/GoLang/solutions"
	"AdventOfCode/AoC2021/GoLang/utils"
	"log"
)

func main() {
	part1 := solutions.D3P1
	part2 := solutions.D3P2

	p1, err := part1()
	if err != nil {
		log.Fatal(err)
	}

	p2, err := part2()
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint([]interface{}{p1, p2})
}
