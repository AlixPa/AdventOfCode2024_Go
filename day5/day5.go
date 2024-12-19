package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var m_rev map[int]map[int]struct{} = make(map[int]map[int]struct{})
var s_pages [][]int = make([][]int, 0)

func extractIntsRegMultiple(s string, s_re string) [][]int {
	sl_ret := make([][]int, 0)
	re := regexp.MustCompile(s_re)
	matchs := re.FindAllStringSubmatch(s, -1)
	for _, match := range matchs {
		s_tmp := make([]int, 0)
		for _, v_s := range match[1:] {
			v, _ := strconv.Atoi(v_s)
			s_tmp = append(s_tmp, v)
		}
		sl_ret = append(sl_ret, s_tmp)
	}
	return sl_ret
}

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
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		if matchsInt := extractIntsRegMultiple(scan.Text(), `(\d+)\|(\d+)`); len(matchsInt) > 0 {
			if _, exists_b := m_rev[matchsInt[0][1]]; !exists_b {
				m_rev[matchsInt[0][1]] = make(map[int]struct{})
			}
			m_rev[matchsInt[0][1]][matchsInt[0][0]] = struct{}{}
		} else if matchsInt := extractIntsRegUnique(scan.Text()); len(matchsInt) > 0 {
			s_pages = append(s_pages, matchsInt)
		}
	}
}

func checkPagesCorrect(left, right int) bool {
	_, corrupted := m_rev[left][right]
	return !corrupted
}

func checkSeqCorrect(seq []int) (bool, int, int) {
	for left_i, left_v := range seq[:len(seq)-1] {
		for shift, right_v := range seq[left_i+1:] {
			if !checkPagesCorrect(left_v, right_v) {
				return false, left_i, left_i + shift + 1
			}
		}
	}
	return true, -1, -1
}

func F1() {
	sum := 0
	for _, seq := range s_pages {
		if o, _, _ := checkSeqCorrect(seq); !o {
			continue
		}
		sum += seq[len(seq)/2]
	}
	fmt.Println(sum)
}

func reorderSeq(seq []int) {
	if o, l, r := checkSeqCorrect(seq); !o {
		seq[l], seq[r] = seq[r], seq[l]
		reorderSeq(seq)
	}
}

func F2() {
	sum := 0
	for _, seq := range s_pages {
		if o, _, _ := checkSeqCorrect(seq); !o {
			reorderSeq(seq)
			sum += seq[len(seq)/2]
		}
	}
	fmt.Println(sum)
}
