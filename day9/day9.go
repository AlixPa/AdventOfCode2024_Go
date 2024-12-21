package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gammazero/deque"
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

var extracted []int
var total_blocs int = 0

func init() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(file)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	scan.Scan()
	extracted = extractIntsRegUnique(scan.Text(), `(\d)`)
	for _, n := range extracted {
		total_blocs += n
	}
}

func F1() {
	sum := 0
	left, right := 0, total_blocs-1
	left_i, right_i := 0, len(extracted)-1
	left_cur, right_cur := 0, 0
	for left <= right {
		if left_cur >= extracted[left_i] {
			left_cur = 0
			left_i++
			continue
		}
		if right_cur >= extracted[right_i] {
			right_cur = 0
			right -= extracted[right_i-1]
			right_i -= 2
			continue
		}
		if left_i%2 == 0 {
			sum += (left_i / 2) * left
		} else {
			sum += (right_i / 2) * left
			right--
			right_cur++
		}
		left++
		left_cur++
	}
	fmt.Println(sum)
}

func insort(m map[int]*deque.Deque[int], key, val int) {
	for i := 0; i < m[key].Len(); i++ {
		if m[key].At(i) > val {
			m[key].Insert(i, val)
			return
		}
	}
	m[key].PushBack(val)
}

func F2() {
	files_chain, true_pos := make([]int, 0), make([]int, 0)
	empty_pos := make(map[int]*deque.Deque[int])
	for i := 0; i <= 9; i++ {
		empty_pos[i] = new(deque.Deque[int])
	}
	cur_pos := 0
	for i, n := range extracted {
		true_pos = append(true_pos, cur_pos)
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				files_chain = append(files_chain, i/2)
			}
		}
		if i%2 == 1 {
			for j := 0; j < n; j++ {
				files_chain = append(files_chain, 0)
			}
			empty_pos[n].PushBack(cur_pos)
		}
		cur_pos += n
	}
	// I now have the corresponding memory state in files_chain
	// And the map from size of free space to its position in ascending order

	for i := len(extracted) - 1; i >= 0; i -= 2 {
		length_to_insert, pos_from := extracted[i], true_pos[i]
		free_space_pos, free_space_size := pos_from, 0
		for free_space_of := length_to_insert; free_space_of <= 9; free_space_of++ {
			if empty_pos[free_space_of].Front() < free_space_pos {
				free_space_pos = empty_pos[free_space_of].Front()
				free_space_size = free_space_of
			}
		}
		if free_space_size == 0 {
			continue
		}
		free_space_pos = empty_pos[free_space_size].PopFront()
		for inserting := 0; inserting < length_to_insert; inserting++ {
			files_chain[free_space_pos+inserting] = i / 2
			files_chain[pos_from+inserting] = 0
		}
		resting_space_size := free_space_size - length_to_insert
		resting_space_pos := free_space_pos + length_to_insert
		insort(empty_pos, resting_space_size, resting_space_pos)
	}

	sum := 0
	for i, id_file := range files_chain {
		sum += i * id_file
	}
	fmt.Println(sum)
}
