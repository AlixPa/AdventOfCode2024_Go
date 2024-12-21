package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

var grid [][]int = make([][]int, 0)

func init() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		grid = append(grid, extractIntsRegUnique(scan.Text(), `(\d)`))
	}
}

var movesNonDiag [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func isInGrid[T any](g [][]T, i, j int) bool {
	return !(i < 0 || j < 0 || i >= len(g) || j >= len(g[i]))
}

func explore1(start_i, start_j, num int, m_seen map[int]struct{}) {
	if num == 9 {
		m_seen[start_i*2*len(grid)+start_j] = struct{}{}
		return
	}
	for _, move := range movesNonDiag {
		if isInGrid(grid, start_i+move[0], start_j+move[1]) && grid[start_i+move[0]][start_j+move[1]] == num+1 {
			explore1(start_i+move[0], start_j+move[1], num+1, m_seen)
		}
	}
}

func F1() {
	sum := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				m_seen := make(map[int]struct{})
				explore1(i, j, 0, m_seen)
				sum += len(m_seen)
			}
		}
	}
	fmt.Println(sum)
}

func explore2(start_i, start_j, num int) int {
	if num == 9 {
		return 1
	}
	ret := 0
	for _, move := range movesNonDiag {
		if isInGrid(grid, start_i+move[0], start_j+move[1]) && grid[start_i+move[0]][start_j+move[1]] == num+1 {
			ret += explore2(start_i+move[0], start_j+move[1], num+1)
		}
	}
	return ret
}

func F2() {
	sum := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				sum += explore2(i, j, 0)
			}
		}
	}
	fmt.Println(sum)
}
