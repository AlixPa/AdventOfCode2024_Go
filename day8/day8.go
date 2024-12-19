package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type position struct {
	i, j int
}

var table []string = make([]string, 0)
var antennas map[rune][]position = make(map[rune][]position)

func init() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		table = append(table, scan.Text())
	}

	for i, line := range table {
		for j, c := range line {
			if c == '.' {
				continue
			}
			if _, o := antennas[c]; !o {
				antennas[c] = make([]position, 0)
			}
			antennas[c] = append(antennas[c], position{i, j})
		}
	}
}

func isOutOfMap(i, j int) bool {
	if i < 0 || j < 0 || i >= len(table) || j >= len(table[i]) {
		return true
	}
	return false
}

func F1() {
	sum := 0
	is_antinode := make([][]bool, len(table))
	var antinode_i, antinode_j int
	for i := range is_antinode {
		is_antinode[i] = make([]bool, len(table[i]))
	}

	for c := range antennas {
		for i, antenna1 := range antennas[c] {
			for _, antenna2 := range antennas[c][i+1:] {
				antinode_i = 2*antenna1.i - antenna2.i
				antinode_j = 2*antenna1.j - antenna2.j
				if !isOutOfMap(antinode_i, antinode_j) && !is_antinode[antinode_i][antinode_j] {
					is_antinode[antinode_i][antinode_j] = true
					sum++
				}
				antinode_i = 2*antenna2.i - antenna1.i
				antinode_j = 2*antenna2.j - antenna1.j
				if !isOutOfMap(antinode_i, antinode_j) && !is_antinode[antinode_i][antinode_j] {
					is_antinode[antinode_i][antinode_j] = true
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}

func F2() {
	sum := 0
	is_antinode := make([][]bool, len(table))
	var antinode_i, antinode_j int
	for i := range is_antinode {
		is_antinode[i] = make([]bool, len(table[i]))
	}

	for c := range antennas {
		for i, antenna1 := range antennas[c] {
			for _, antenna2 := range antennas[c][i+1:] {
				antinode_i = antenna1.i
				antinode_j = antenna1.j
				for {
					if isOutOfMap(antinode_i, antinode_j) {
						break
					}
					if !is_antinode[antinode_i][antinode_j] {
						is_antinode[antinode_i][antinode_j] = true
						sum++
					}
					antinode_i += antenna1.i - antenna2.i
					antinode_j += antenna1.j - antenna2.j
				}
				antinode_i = antenna2.i
				antinode_j = antenna2.j
				for {
					if isOutOfMap(antinode_i, antinode_j) {
						break
					}
					if !is_antinode[antinode_i][antinode_j] {
						is_antinode[antinode_i][antinode_j] = true
						sum++
					}
					antinode_i -= antenna1.i - antenna2.i
					antinode_j -= antenna1.j - antenna2.j
				}
			}
		}
	}
	fmt.Println(sum)
}
