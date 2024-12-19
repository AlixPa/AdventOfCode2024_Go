package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var table []string

func init() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		table = append(table, scan.Text())
	}
}

func checkXMAS(move [2]int, i, j int) bool {
	if table[i][j] != 'X' {
		return false
	}
	if table[i+move[0]][j+move[1]] != 'M' {
		return false
	}
	if table[i+2*move[0]][j+2*move[1]] != 'A' {
		return false
	}
	if table[i+3*move[0]][j+3*move[1]] != 'S' {
		return false
	}
	return true
}

func checkNum(move [2]int) int {
	sum := 0
	for i := range table {
		for j := range table[i] {
			if (i+3*move[0]) >= len(table) || (j+3*move[1]) >= len(table[i]) || (i+3*move[0]) < 0 || (j+3*move[1]) < 0 {
				continue
			}
			if checkXMAS(move, i, j) {
				sum += 1
			}
		}
	}
	return sum
}

func F1() {
	var movesDiag [8][2]int = [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

	sum := 0
	for _, v := range movesDiag {
		sum += checkNum(v)
	}
	fmt.Println(sum)
}

func checkDiag(move [2]int, i, j int) bool {
	return (table[i-move[0]][j-move[1]] == 'M' && table[i+move[0]][j+move[1]] == 'S') || (table[i-move[0]][j-move[1]] == 'S' && table[i+move[0]][j+move[1]] == 'M')
}

func checkMAS(i, j int) bool {
	return checkDiag([2]int{1, 1}, i, j) && checkDiag([2]int{-1, 1}, i, j)
}

func F2() {
	sum := 0
	for i := range table {
		for j := range table {
			if i == 0 || i == len(table)-1 || j == 0 || j == len(table[i])-1 {
				continue
			}
			if table[i][j] == 'A' && checkMAS(i, j) {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}
