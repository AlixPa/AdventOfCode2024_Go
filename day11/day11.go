package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var stones []int = make([]int, 0)
var digitBlinks [10][76]int

/*
	Returns the list of integers captured by the s_re regexp expression in the string s

No regex expression will just extract the integers seen in s.
*/
func extractIntsRegUnique(s string, s_re ...string) []int {
	sl_ret := make([]int, 0)
	if len(s_re) == 0 {
		s_re = []string{`(\d+)`}
	}
	re := regexp.MustCompile(s_re[0])
	matchs := re.FindAllStringSubmatch(s, -1)
	for _, match := range matchs {
		v, _ := strconv.Atoi(match[1])
		sl_ret = append(sl_ret, v)
	}
	return sl_ret
}

func init() {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	scan.Scan()
	stones = extractIntsRegUnique(scan.Text())
	initDigitBlinks()
}

func nbDigitInt(n int) int {
	nb := 0
	for n != 0 {
		nb++
		n /= 10
	}
	return nb
}

/* Returns base^pow */
func powInt(base, pow int) int {
	res := 1
	if pow <= 0 {
		return res
	}
	for {
		if (pow & 1) != 0 {
			res *= base
		}
		pow >>= 1
		if pow == 0 {
			return res
		}
		base *= base
	}
}

func getNumberStones(stone, steps int) int {
	if steps <= 0 {
		return 1
	}
	if stone <= 9 {
		return digitBlinks[stone][steps]
	}
	if nbDigitInt(stone)%2 == 0 {
		return getNumberStones(stone/powInt(10, nbDigitInt(stone)/2), steps-1) + getNumberStones(stone%powInt(10, nbDigitInt(stone)/2), steps-1)
	} else {
		return getNumberStones(stone*2024, steps-1)
	}
}

func initDigitBlinks() {
	for i := 0; i <= 9; i++ {
		digitBlinks[i][0] = 1
	}
	for i := 1; i < len(digitBlinks[0]); i++ {
		digitBlinks[0][i] = digitBlinks[1][i-1]
		for j := 1; j <= 9; j++ {
			digitBlinks[j][i] = getNumberStones(2024*j, i-1)
		}
	}
}

func countSteps(steps int) int {
	sum := 0
	for _, stone := range stones {
		sum += getNumberStones(stone, steps)
	}
	return sum
}

func F1() {
	fmt.Println(countSteps(25))
}

func F2() {
	fmt.Println(countSteps(75))
}
