package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func IntToIntSlice(num int) (res []int) {
	numStr := strconv.Itoa(num)
	digits := strings.Split(numStr, "")
	for _, d := range digits {
		n, _ := strconv.Atoi(d)
		res = append(res, n)
	}
	return res
}

func MatchCounts(passDigits []int) []int {
	res := []int{}
	currentCount := 0
	for i, d := range passDigits {
		if i > 0 {
			if passDigits[i-1] == d {
				currentCount++
			} else {
				if currentCount > 0 {
					res = append(res, currentCount + 1)
				}
				currentCount = 0
			}
		}
	}
	if currentCount > 0 {
		res = append(res, currentCount + 1)
	}
	return res
}

func IsValid(password int) bool {
	adjecentFlag := false
	digits := IntToIntSlice(password)
	lastD := 0
	for i, d := range digits {
		if i > 0 {
			if lastD == d {
				adjecentFlag = true
			}
			if d < lastD {
				return false
			}
		}
		lastD = d
	}
	return adjecentFlag == true
}

func IsValid2(password int) bool {
	digits := IntToIntSlice(password)
	lastD := 0
	for i, d := range digits {
		if i > 0 {
			if d < lastD {
				return false
			}
		}
		lastD = d
	}
	matchCounts := MatchCounts(digits)
	for _, m := range matchCounts {
		if m == 2 {
			return true
		}
	}
	return false
}

func GetValidCountInRange(min, max, version int) (validCount int) {
	if version == 1 {
		for pass := min; pass <= max; pass++ {
			if IsValid(pass) {
				validCount++
			}
		}
	}
	if version == 2 {
		for pass := min; pass <= max; pass++ {
			if IsValid2(pass) {
				validCount++
			}
		}
	}
	return validCount
}

func DoSilver() {
	fmt.Println("Solution: ", GetValidCountInRange(197487, 673251, 1))
}

func DoGold() {
	fmt.Println("Solution: ", GetValidCountInRange(197487, 673251, 2))
}