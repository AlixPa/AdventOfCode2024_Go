package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var grid [][]rune = make([][]rune, 0)

func init() {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		tmp := make([]rune, len(scan.Text()))
		for i, b := range scan.Text() {
			tmp[i] = b
		}
		grid = append(grid, tmp)
	}
}

func isInGrid[T any](g [][]T, i, j int) bool {
	return !(i < 0 || j < 0 || i >= len(g) || j >= len(g[i]))
}

var movesNonDiag [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func exploreZonePerim(i, j int, c rune, seen [][]bool, area, perim *int) {
	if !isInGrid(grid, i, j) || grid[i][j] != c || seen[i][j] {
		return
	}
	seen[i][j] = true
	*area++
	for _, move := range movesNonDiag {
		if !isInGrid(grid, i+move[0], j+move[1]) || grid[i+move[0]][j+move[1]] != c {
			*perim++
		}
	}
	for _, move := range movesNonDiag {
		exploreZonePerim(i+move[0], j+move[1], c, seen, area, perim)
	}
}

func count(f func(i, j int, c rune, seen [][]bool, area, perim *int)) int {
	seen := make([][]bool, 0)
	for i := range grid {
		seen = append(seen, make([]bool, len(grid[i])))
	}

	sum := 0
	for i := range grid {
		for j := range grid[i] {
			area, perim_or_sides := 0, 0
			f(i, j, grid[i][j], seen, &area, &perim_or_sides)
			sum += area * perim_or_sides
		}
	}
	return sum
}

func F1() {
	fmt.Println(count(exploreZonePerim))
}

func isNotC(i, j int, c rune) bool {
	return !isInGrid(grid, i, j) || grid[i][j] != c
}

func numberInnerAngle(i, j int, c rune) int {
	nb := 0
	for k := range movesNonDiag {
		if !isNotC(i+movesNonDiag[k][0], j+movesNonDiag[k][1], c) && !isNotC(i+movesNonDiag[(k+1)%4][0], j+movesNonDiag[(k+1)%4][1], c) && isNotC(i+movesNonDiag[k][0]+movesNonDiag[(k+1)%4][0], j+movesNonDiag[k][1]+movesNonDiag[(k+1)%4][1], c) {
			nb++
		}
	}
	return nb
}

func numberOuterAngle(i, j int, c rune) int {
	nb := 0
	for k := range movesNonDiag {
		if isNotC(i+movesNonDiag[k][0], j+movesNonDiag[k][1], c) && isNotC(i+movesNonDiag[(k+1)%4][0], j+movesNonDiag[(k+1)%4][1], c) {
			nb++
		}
	}
	return nb
}

func exploreZoneSides(i, j int, c rune, seen [][]bool, area, sides *int) {
	if !isInGrid(grid, i, j) || grid[i][j] != c || seen[i][j] {
		return
	}
	seen[i][j] = true
	*area++
	*sides += numberInnerAngle(i, j, c)
	*sides += numberOuterAngle(i, j, c)

	for _, move := range movesNonDiag {
		exploreZoneSides(i+move[0], j+move[1], c, seen, area, sides)
	}
}

func F2() {
	fmt.Println(count(exploreZoneSides))
}
