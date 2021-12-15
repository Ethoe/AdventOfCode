package solutions

import (
	"AdventOfCode/AoC2021/GoLang/utils"
	"fmt"
	"strconv"
)

func D3P1() (int, error) {
	text, err := utils.ReadFile("day3.txt")
	if err != nil {
		return 0, err
	}
	gamma := make(map[int]int)
	for _, v := range text {
		for num := 0; num < len(v); num++ {
			fmt.Println(string(v[num]))
			if string(v[num]) == "1" {
				gamma[num] += 1
			} else {
				print(gamma[num])

				gamma[num] -= 1
			}
		}
	}
	g := ""
	e := ""
	for num := 0; num < len(gamma); num++ {
		if gamma[num] > 0 {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	gam, _ := strconv.ParseInt(g, 2, 64)
	eps, _ := strconv.ParseInt(e, 2, 64)

	return int(gam * eps), nil
}

func D3P2() (int, error) {
	text, err := utils.ReadFile("day3.txt")
	if err != nil {
		return 0, err
	}

	oxygen := lifeSupportRating(text, 0, true)
	co2 := lifeSupportRating(text, 0, false)
	return oxygen * co2, nil
}

func lifeSupportRating(numbers []string, index int, most bool) int {
	ones := 0
	zeros := 0
	positive := make([]string, 0)
	negative := make([]string, 0)

	for _, v := range numbers {
		if string(v[index]) == "1" {
			ones += 1
			positive = append(positive, v)
		} else {
			zeros += 1
			negative = append(negative, v)
		}
	}

	var newNumbers []string

	if most {
		if ones >= zeros {
			newNumbers = positive
		} else {
			newNumbers = negative
		}
	} else {
		if zeros <= ones {
			newNumbers = negative
		} else {
			newNumbers = positive
		}
	}

	if len(newNumbers) == 1 {
		res, _ := strconv.ParseInt(newNumbers[0], 2, 64)
		return int(res)
	}

	return lifeSupportRating(newNumbers, index+1, most)
}
