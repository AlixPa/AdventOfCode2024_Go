package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

var table [][]int = make([][]int, 0)

func init() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		table = append(table, extractIntsRegUnique(scan.Text()))
	}
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

func F1() {
	sum := 0

	for _, line := range table {
		for combi := range powInt(2, len(line)-2) {
			res := line[1]
			for _, next_num := range line[2:] {
				if combi%2 == 0 {
					res += next_num
				} else {
					res *= next_num
				}
				combi /= 2
			}
			if res == line[0] {
				sum += res
				break
			}
		}
	}
	fmt.Println(sum)
}

func combineInts(x, y int) int {
	ret, _ := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
	return ret
}

func F2() {
	sum := 0

	for _, line := range table {
		for combi := range powInt(3, len(line)-2) {
			res := line[1]
			for _, next_num := range line[2:] {
				if combi%3 == 0 {
					res += next_num
				} else if combi%3 == 1 {
					res *= next_num
				} else {
					res = combineInts(res, next_num)
				}
				combi /= 3
			}
			if res == line[0] {
				sum += res
				break
			}
		}
	}
	fmt.Println(sum)
}
