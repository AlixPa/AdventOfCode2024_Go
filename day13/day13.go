package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var buttonsA, buttonsB, objectives [][]int = make([][]int, 0), make([][]int, 0), make([][]int, 0)

/* Returns the list of the captured grouped of each match of the s_re regexp expression in the string s. */
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

func init() {
	file, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for i := 0; scan.Scan(); i++ {
		i %= 4
		switch i {
		case 0:
			buttonsA = append(buttonsA, extractIntsRegMultiple(scan.Text(), `X\+(\d+), Y\+(\d+)`)[0])
		case 1:
			buttonsB = append(buttonsB, extractIntsRegMultiple(scan.Text(), `X\+(\d+), Y\+(\d+)`)[0])
		case 2:
			objectives = append(objectives, extractIntsRegMultiple(scan.Text(), `X=(\d+), Y=(\d+)`)[0])
		}
	}

}

func isWinner(machine, combi int) bool {
	return objectives[machine][0] == buttonsA[machine][0]*(combi/101)+buttonsB[machine][0]*(combi%101) && objectives[machine][1] == buttonsA[machine][1]*(combi/101)+buttonsB[machine][1]*(combi%101)
}

func count(cdt func(n, m int) bool) int {
	sum := 0
	for i := range buttonsA {
		/*
			Ax*Na + Bx*Nb = Ox
			Ay*Na + By*Nb = Oy

			We search Na, Nb
			Ax*Na = Ox - Bx*Nb
			Na = (Ox - Bx*Nb) / Ax = Ox/Ax - Bx*Nb/Ax

			(Ay*(Ox - Bx*Nb) / Ax) + By*Nb = Oy
			(Ay*(Ox - Bx*Nb) / Ax) + By*Nb*Ax/Ax = Oy
			[Ay*(Ox - Bx*Nb) + By*Nb*Ax] / Ax = Oy
			Ay*(Ox - Bx*Nb) + By*Nb*Ax = Oy*Ax
			Ay*Ox - Ay*Bx*Nb + By*Nb*Ax = Oy*Ax
			Nb*(Ax*By - Ay*Bx) = Oy*Ax - Ox*Ay

			Nb = (Oy*Ax - Ox*Ay) / (Ax*By - Ay*Bx)
			Na = (Ox - Bx*Nb) / Ax
		*/
		denom := (buttonsA[i][0]*buttonsB[i][1] - buttonsA[i][1]*buttonsB[i][0])
		num := (objectives[i][1]*buttonsA[i][0] - objectives[i][0]*buttonsA[i][1])
		if denom == 0 || (num/denom)*denom != num {
			continue
		}
		nb := num / denom
		num = objectives[i][0] - buttonsB[i][0]*nb
		denom = buttonsA[i][0]
		if (num/denom)*denom != num {
			continue
		}
		na := num / denom
		if cdt(na, nb) {
			sum += 3*na + nb
		}
	}
	return sum
}

func F1() {
	fmt.Println(count(func(n, m int) bool { return n >= 0 && m >= 0 && n <= 100 && m <= 100 }))
}

func F2() {
	to_add := 10000000000000
	for _, obj := range objectives {
		obj[0], obj[1] = obj[0]+to_add, obj[1]+to_add
	}
	fmt.Println(count(func(n, m int) bool { return n >= 0 && m >= 0 }))
}
