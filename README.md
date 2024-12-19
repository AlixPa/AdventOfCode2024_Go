# Advent of code 2024 in Go

This is the start of my Go learning journey.

This repo contains all my code produced for the Advent of Code 2024. As it was my first time using Go, the first days may not use the most precise functionnalities of Go, nor may be efficient.

There will be a small comment on each day resolution, as well as a complexity analysis and futur amelioration ideas (implemented for some).

## Day 1

Input is two lists of n numbers.

### Task 1

Solved using sorting on slices.

Each list of location being sorted,
we then substract one to the other for each line.

Implementation of a few usefull functions.
> "func absInt(x int) int" has been added to the snippets

Complexity is O(n*log(n)).

### Task 2

Solved using counting in a map.

We first count the occurences of each location of the second list,
then multiply it by the location if present in the first list.

Complexity is O(n).

## Day 2

Input is a list of n lists of p numbers (or less).

### Task 1

Solved using verification of the ordering.

We first check what is the ordering (going up or down) of the list,
then check if each element respect the ordering it should.

Complexity is O(n*p).

### Task 2

Solved iterating over the functions of task 1.

For each number of a report,
we check if the report is correct without the number.
> "func numComp[T cmp.Ordered](a, b T) int" has been added to snippets. Allows to sort []int, []float, etc with slices.SortFunc.

Complexity is O(n*p<sup>2</sup>).

## Day 3

Input is a string of size n.

### Task 1

Solved using regex to match the given pattern.

We simply extract each "mul(x,y)" using the regex expression `mul\((\d+),(\d+)\)`.

Complexity is O(n).

### Task 2

Solved also using regex to get the active patterns.

We get the index of each do() and don't(),
then only use the mul(x,y) if they are after a do() but before any don't().
> "type intBool struct" has been added to the snippet. It represents the tuple (int, bool).
> "func compIntBool(ib1, ib2 intBool) int" has been added to the snippet. Allows to sort [](int, bool) with slices.SortFunc.

Complexity is O(n).

## Day 4

Input is a grid of characters of n lines and p columns.

### Task 1

Solved using local exploration.

For each position in the grid of characters,
we check if it is the start of an "XMAS" in one or another direction.
> "var movesDiag [8][2]int" has been added to snippets. It is the list of possibles movements of 1 step on a grid.
> "var movesNonDiag [8][2]int" has been added to snippets. Same as the above, without the diagonal movements.

Complexity is O(n*p).

### Task 2

Solved also using local exploration.

For each position in the grid of characters,
we check if it is an "A", if yes we check if "M.S" are written on each diagonal.

Complexity is O(n*p).

## Day 5

Input is a list of n numbers, then a list of m lists of p numbers.
> "func extractIntsRegMultiple(s string, s_re string) [][]int" has been added to snippets. Extracts groups of integers from a string using a regex expression with multiple capturing groups.
> "func extractIntsRegUnique(s string, s_re ...string) []int" has been added to snippets. Extracts integers from a string using a regex expression with unique capturing group.

### Task 1

Solved using map on reversing order.

For each rule p1 | p2, we store p1 in map[p2],
then for each sequence of pages, we check if the pages are all ordered correctly,

Complexity is O(n + m*p<sup>2</sup>).

### Task 2

Solved using an unoptimised sorting technique.

To reorder the sequences of pages, we used kindof an unoptimized bubble sort.

Complexity is O(n + m*p<sup>4</sup>). Could easily reduce one p factor by using real bubble sort.

## Day 6

Input is a grid of characters of n lines and p columns.

### Task 1

Solved using the grid exploration.

We made the gard move until it exits the map, counting new position he has been on.
(This suppose the gard can exits the map).

Complexity is 0(n*p).

### Task 2

Solved using multiple grid exploration.

For each position of the grid, we add an obstacle on it,
then we make the gard explore the map and check if it loops or not.

Complexity is 0(n<sup>2</sup>*p<sup>2</sup>).

## Day 7

Input is a list of n lists of p numbers.

### Task 1

Solved using combinaisons.

For each combinaison of "+" and "*", we calculate the result and look if it matchs the expected result.
> "func powInt(base, pow int) int" has been added to snippets. Calculate base^pow with the "exponentiation by squaring" technique.

Complexity is O(n\*2<sup>p</sup>\*p). (Note that p is supper small).

### Task 2

Solved using combinaisons.

Same as the above, adding the "||" operation as a function.

Complexity is O(n\*3<sup>p</sup>\*p). (Note that p is supper small).

## Day 8

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 9

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 10

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 11

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 12

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 13

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 14

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 15

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 16

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 17

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 18

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 19

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 20

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 21

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 22

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 23

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is

## Day 24

Input is

### Task 1

Solved using combinaisons.



Complexity is

### Task 2

Solved using



Complexity is