package solutions

import (
	"AdventOfCode/AoC2020/GoLang/utils"
	"errors"
	"strconv"
)

/*
After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island.
Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of
a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow,
you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar;
the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input);
apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
In this list, the two entries that sum to 2020 are 1721 and 299.
Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.
*/
func D1P1() (int, error) {
	text, err := utils.ReadFile("day1.txt")
	if err != nil {
		return 0, err
	}

	hash := make(map[int]int, 0)
	for _, item := range text {
		num, err := strconv.Atoi(item)
		if err != nil {
			return 0, err
		}
		if hash[2020-num] == 0 {
			hash[num] = num
		} else {
			return num * hash[2020-num], nil
		}
	}
	return 0, errors.New("no match found")
}

func D1P2() (int, error) {
	text, err := utils.ReadFile("day1.txt")
	if err != nil {
		return 0, err
	}

	items := make([]int, 0)
	for _, item := range text {
		num, err := strconv.Atoi(item)
		if err != nil {
			return 0, err
		}
		items = append(items, num)
	}

	for _, a := range items {
		for _, b := range items {
			for _, c := range items {
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}
	return 0, errors.New("could not find solution")
}
