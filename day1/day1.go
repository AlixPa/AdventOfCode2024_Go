package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func absInt(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func compInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

var tab1, tab2 []int

func init() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var v1, v2 int
	re := regexp.MustCompile(`(\d+)`)
	var match [][]string
	for fileScanner.Scan() {
		match = re.FindAllStringSubmatch(fileScanner.Text(), -1)
		v1, _ = strconv.Atoi(match[0][1])
		v2, _ = strconv.Atoi(match[1][1])
		tab1 = append(tab1, v1)
		tab2 = append(tab2, v2)
	}
	slices.SortFunc(tab1, compInt)
	slices.SortFunc(tab2, compInt)
}

func F1() {
	var sum int = 0
	for i := 0; i < len(tab1); i++ {
		sum += absInt(tab1[i], tab2[i])
	}
	fmt.Println(sum)
}

func F2() {
	var m_nums map[int]int = make(map[int]int)
	for i := range tab1 {
		m_nums[tab1[i]] = 0
		m_nums[tab2[i]] = 0
	}

	for _, n := range tab2 {
		m_nums[n] += 1
	}
	sum := 0
	for _, n := range tab1 {
		sum += n * m_nums[n]
	}
	fmt.Println(sum)
}
