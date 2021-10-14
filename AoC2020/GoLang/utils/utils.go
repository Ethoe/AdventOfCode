package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(file string) (res []string, err error) {
	res = make([]string, 0)

	f, err := os.Open("AoC2020\\files\\" + file)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
}

func PrettyPrint(items []interface{}) {
	for i, v := range items {
		fmt.Printf("Part %v: %v\n", i+1, v)
	}
}
