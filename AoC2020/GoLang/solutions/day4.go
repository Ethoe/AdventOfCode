package solutions

import (
	"AdventOfCode/AoC2020/GoLang/utils"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	fields map[string]string
}

func buildPassport(text []string, index *int) (passport, error) {
	if *index >= len(text) {
		return passport{}, errors.New("out of bounds")
	}
	res := passport{make(map[string]string, 0)}
	for true {
		if *index >= len(text) || text[*index] == "" {
			*index++
			return res, nil
		} else {
			tokens := strings.Fields(text[*index])
			for _, token := range tokens {
				field := strings.Split(token, ":")
				res.fields[field[0]] = field[1]
			}
			*index++
		}
	}
	return passport{}, errors.New("oh god")
}

func D4P1() (total int, err error) {
	text, err := utils.ReadFile("day4.txt")
	if err != nil {
		return 0, err
	}

	total = 0
	index := 0

	for index < len(text) {
		var pass passport
		pass, err = buildPassport(text, &index)
		if err != nil {
			return 0, err
		}
		add := true
		for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if _, ok := pass.fields[field]; !ok {
				add = false
			}
		}
		if add {
			total++
		}
	}
	return
}

func colorValidate(hcl string) bool {
	match, err := regexp.MatchString(`#[0-9a-f]{6}`, hcl)
	if err != nil {
		return false
	}
	return match
}

func id(pid string) bool {
	match, err := regexp.MatchString(`[0-9]{9}`, pid)
	if err != nil {
		return false
	}
	return match
}

func contains(test string) bool {
	for _, match := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if test == match {
			return true
		}
	}
	return false
}

func validate(pass passport) bool {
	if byr, err := strconv.Atoi(pass.fields["byr"]); err != nil || byr > 2002 || byr <= 1920 {
		return false
	}
	if iyr, err := strconv.Atoi(pass.fields["iyr"]); err != nil || iyr > 2020 || iyr < 2010 {
		return false
	}
	if eyr, err := strconv.Atoi(pass.fields["eyr"]); err != nil || eyr > 2030 || eyr < 2020 {
		return false
	}
	if len(pass.fields["hgt"]) >= 2 {
		if unit := pass.fields["hgt"][len(pass.fields["hgt"])-2:]; unit == "cm" || unit == "in" {
			if hgt, err := strconv.Atoi(pass.fields["hgt"][:len(pass.fields["hgt"])-2]); err != nil {
				return false
			} else if unit == "cm" {
				if hgt > 193 || hgt <= 150 {
					return false
				}
			} else if hgt >= 76 || hgt <= 59 {
				return false
			}
		}
	}
	if !colorValidate(pass.fields["hcl"]) {
		return false
	}
	if !contains(pass.fields["ecl"]) {
		return false
	}
	if !id(pass.fields["pid"]) {
		return false
	}
	return true
}

func D4P2() (total int, err error) {
	text, err := utils.ReadFile("day4.txt")
	if err != nil {
		return 0, err
	}

	total = 0
	index := 0

	for index < len(text) {
		var pass passport
		pass, err = buildPassport(text, &index)
		if err != nil {
			return 0, err
		}
		if validate(pass) {
			total++
		}
	}
	return
}
