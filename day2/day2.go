package day2

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func numComp[T cmp.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

var table [][]int

func init() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	re := regexp.MustCompile(`(\d+)`)
	var match [][]string
	var v int
	for scan.Scan() {
		match = re.FindAllStringSubmatch(scan.Text(), -1)
		tmp_s := make([]int, 0)
		for _, nS := range match {
			v, _ = strconv.Atoi(nS[1])
			tmp_s = append(tmp_s, v)
		}
		table = append(table, tmp_s)
	}
}

func checkIfCorrect(s []int) bool {
	order := numComp(s[0], s[1])
	for i := range len(s) - 1 {
		if order != numComp(s[i], s[i+1]) || intAbs(s[i]-s[i+1]) < 1 || intAbs(s[i]-s[i+1]) > 3 {
			return false
		}
	}
	return true
}

func F1() {
	var sum int = 0
	for _, tS := range table {
		if checkIfCorrect(tS) {
			sum++
		}
	}
	fmt.Println(sum)
}

func F2() {
	var sum int = 0
	var tmp_s []int = make([]int, 20)
	for _, tS := range table {
		for i := range len(tS) {
			copy(tmp_s, tS[:i])
			copy(tmp_s[i:], tS[i+1:])
			if checkIfCorrect(tmp_s[:len(tS)-1]) {
				sum++
				break
			}
		}
	}
	fmt.Println(sum)
}
