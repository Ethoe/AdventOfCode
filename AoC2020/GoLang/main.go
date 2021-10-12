package main

import (
	"AdventOfCode/AoC2020/GoLang/solutions"
	"AdventOfCode/AoC2020/GoLang/utils"
	"log"
)

func main() {
	p1, err := solutions.D1P1()
	if err != nil {
		log.Fatal(err)
	}

	p2, err := solutions.D1P2()
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint([]interface{}{p1, p2})
}
