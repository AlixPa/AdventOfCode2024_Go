package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var table []string = make([]string, 0)

func init() {
	file, err := os.Open("day6/input.txt")
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

func initSeenSart() (seen [][]bool, start_i int, start_j int) {
	seen = make([][]bool, 0)
	for i := range table {
		seen_line := make([]bool, len(table[i]))
		for j := range table[i] {
			if table[i][j] == '^' {
				start_i, start_j = i, j
			}
		}
		seen = append(seen, seen_line)
	}
	return
}

func outOfMap(pos_i, pos_j int) bool {
	if pos_i < 0 || pos_j < 0 {
		return true
	}
	if pos_i >= len(table) || pos_j >= len(table[0]) {
		return true
	}
	return false
}

func F1() {
	seen, pos_i, pos_j := initSeenSart()
	var movesNonDiag [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	next_direction := 0
	nb_seen, nb_loop := 0, 0

	for {
		nb_loop++
		if !seen[pos_i][pos_j] {
			nb_seen++
			seen[pos_i][pos_j] = true
		}
		if outOfMap(pos_i+movesNonDiag[next_direction][0], pos_j+movesNonDiag[next_direction][1]) {
			break
		} else if table[pos_i+movesNonDiag[next_direction][0]][pos_j+movesNonDiag[next_direction][1]] == '#' {
			next_direction += 1
			next_direction %= 4
		} else {
			pos_i += movesNonDiag[next_direction][0]
			pos_j += movesNonDiag[next_direction][1]
		}
	}
	fmt.Println(nb_seen)
}

func isLooped(obst_i, obst_j int) bool {
	save := table[obst_i]
	table[obst_i] = table[obst_i][:obst_j] + "#" + table[obst_i][obst_j+1:]
	seen, pos_i, pos_j := initSeenSart()
	var movesNonDiag [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	next_direction := 0
	nb_seen, nb_loop := 0, 0

	for {
		nb_loop++
		if !seen[pos_i][pos_j] {
			nb_seen++
			seen[pos_i][pos_j] = true
			nb_loop = 0
		}
		if outOfMap(pos_i+movesNonDiag[next_direction][0], pos_j+movesNonDiag[next_direction][1]) {
			break
		} else if table[pos_i+movesNonDiag[next_direction][0]][pos_j+movesNonDiag[next_direction][1]] == '#' {
			next_direction += 1
			next_direction %= 4
		} else {
			pos_i += movesNonDiag[next_direction][0]
			pos_j += movesNonDiag[next_direction][1]
		}
		if nb_loop > 4*len(table) {
			break
		}
	}
	table[obst_i] = save
	return nb_loop > 4*len(table)
}

func F2() {
	nb_loop := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == '.' && isLooped(i, j) {
				nb_loop++
			}
		}
	}
	fmt.Println(nb_loop)
}
