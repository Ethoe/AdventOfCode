package solutions

import (
	"AdventOfCode/AoC2020/GoLang/utils"
	"errors"
	"strconv"
	"strings"
)

func D5P1() (int, error) {
	text, err := utils.ReadFile("day5.txt")
	if err != nil {
		return 0, err
	}

	highest := 0

	for _, pass := range text {
		row := pass[:7]
		col := pass[7:]
		rowBin := strings.Replace(strings.Replace(row, "F", "0", -1), "B", "1", -1)
		colBin := strings.Replace(strings.Replace(col, "L", "0", -1), "R", "1", -1)
		rowNum, err := strconv.ParseInt(rowBin, 2, 64)
		if err != nil {
			return 0, err
		}
		colNum, err := strconv.ParseInt(colBin, 2, 64)
		if err != nil {
			return 0, err
		}
		if curr := int(rowNum*8 + colNum); curr > highest {
			highest = curr
		}
	}
	return highest, nil
}

func D5P2() (int, error) {
	text, err := utils.ReadFile("day5.txt")
	if err != nil {
		return 0, err
	}

	seatList := make(map[int]bool, 0)
	highest := 0

	for _, pass := range text {
		row := pass[:7]
		col := pass[7:]
		rowBin := strings.Replace(strings.Replace(row, "F", "0", -1), "B", "1", -1)
		colBin := strings.Replace(strings.Replace(col, "L", "0", -1), "R", "1", -1)
		rowNum, err := strconv.ParseInt(rowBin, 2, 64)
		if err != nil {
			return 0, err
		}
		colNum, err := strconv.ParseInt(colBin, 2, 64)
		if err != nil {
			return 0, err
		}
		if curr := int(rowNum*8 + colNum); curr > highest {
			highest = curr
		}
		seatList[int(rowNum*8+colNum)] = true
	}
	for i := 0; i < highest; i++ {
		if seatList[i-1] && seatList[i+1] && !seatList[i] {
			return i, nil
		}
	}
	return 0, errors.New("does not exist")
}
