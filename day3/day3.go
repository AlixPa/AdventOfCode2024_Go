package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func F1() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var match [][]string
	sum := 0

	for scan.Scan() {
		match = re.FindAllStringSubmatch(scan.Text(), -1)
		for _, s := range match {
			v1, _ := strconv.Atoi(s[1])
			v2, _ := strconv.Atoi(s[2])
			sum += v1 * v2
		}
	}
	fmt.Println(sum)
}

type intBool struct {
	i int
	b bool
}

func compIntBool(ib1, ib2 intBool) int {
	if ib1.i < ib2.i {
		return -1
	} else if ib1.i > ib2.i {
		return 1
	} else {
		return 0
	}
}

func F2() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)
	sum := 0

	for scan.Scan() {
		matchDo := reDo.FindAllStringSubmatchIndex(scan.Text(), -1)
		matchDont := reDont.FindAllStringSubmatchIndex(scan.Text(), -1)
		sDoDont := make([]intBool, 0)
		sDoDont = append(sDoDont, intBool{0, true})
		for _, sD := range matchDo {
			sDoDont = append(sDoDont, intBool{sD[0], true})
		}
		for _, sDn := range matchDont {
			sDoDont = append(sDoDont, intBool{sDn[0], false})
		}
		slices.SortFunc(sDoDont, compIntBool)
		sDoDontClean := make([]intBool, 0)
		last_bool := true
		for _, v := range sDoDont {
			if last_bool != v.b {
				sDoDontClean = append(sDoDontClean, v)
				last_bool = !last_bool
			}
		}
		fmt.Println(sDoDontClean)
		prev := intBool{0, true}
		var match [][]string
		for _, v := range sDoDontClean {
			if !prev.b {
				prev = v
				continue
			}
			match = reMul.FindAllStringSubmatch(scan.Text()[prev.i:v.i+1], -1)
			for _, s := range match {
				v1, _ := strconv.Atoi(s[1])
				v2, _ := strconv.Atoi(s[2])
				sum += v1 * v2
			}
			prev = v
		}
		if prev.b {
			match = reMul.FindAllStringSubmatch(scan.Text()[prev.i:], -1)
			for _, s := range match {
				v1, _ := strconv.Atoi(s[1])
				v2, _ := strconv.Atoi(s[2])
				sum += v1 * v2
			}
		}
	}
	fmt.Println(sum)
}
