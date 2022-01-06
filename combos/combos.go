package combos

import "go-exercises/insertion_sort"



func Combos(numbers []int) [][]int {
	index := find(numbers, func(i int) bool { return i > 1 })
	if index == -1 { return [][]int {numbers} }

	minuend := numbers[index]
	sets := [][]int { numbers }

	for subrahend := 1; subrahend <= (minuend / 2); subrahend++ {
		difference := minuend - subrahend

		newNumbers := remove(numbers, index)
		newSet := insertion_sort.InsertionSort(append(newNumbers, subrahend, difference))
		sets = append(sets, Combos(newSet)...)
	}

	return sets
}

func find(numbers []int, predicate func(i int) bool) int {
	for i := 0; i < len(numbers); i++ {
		number := numbers[i]
		if predicate(number) { return i }
	}
	return -1
}

func remove(slice []int, index int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, slice[:index]...)
	return append(newSlice, slice[index+1:]...)

	/*
	Earlier was implemented as:
	`return append(slice[:index], slice[index+1:]...)`

	However `slice[:index]` returns a new slice that uses the same underlying array.
	This means different slices share the same array, leading to unexpected behaviour due to mutation.
	Creating a new slice using `make` makes sure a new array is created.
	*/
}

/*
1. sort
2. check index

*/

/*


1. find index of first number greater than one
2. if no index return [][]int {n}
3. minuend = n[index]
4. init possible sets []
5. iterate from minuend - 1 to minuend - (minuend - 2) where difference >= 1 |difference| (subtrahend = minuend - difference)
		1. remove minuend from n
		2. newSet = [n..., subrahend, difference]
		3. concat(possible sets, combos(newSet)...)
4. return possible sets


minuend == 1
=> no iterations

minuend == 2
=> one iterations
	subtrahend == 1

miuend == 3
=> one iterations
	subtrahend == 1 2, 1

minuend == 4
=> two iterations
	subtrahend == 1 3, 1
	subtrahend == 2 2, 2

minuend == 5
=> two iterations
	subtrahend == 1 4, 1
	subtrahend == 2 3, 2

subrahend iterates from 1 to (minuend / 2)







                    {5}

      {1, 4}                     {2, 3}

{1, 1, 3} {1, 2, 2}

{1, 1, 1, 2}

{1, 1, 1, 1, 1}


30 > num > 0
positive integers that sum
n = 10 => ~6000ms
sub-array don't need numbers sorted
sub-array don't need sorted in any order

Mental Model:
use recursion to branch possible sets then concat possible sets
use hash index to check if possible branch already computed

Data Structure:
recursion + [][]int
nothing special

Algorithm:

params
n = []int
set_index = {}[]int

1. find index of first number greater than one
2. if no index return [][]int {n}
3. minuend = n[index]
4. init possible sets []
5. iterate from minuend - 1 to minuend - (minuend - 2) where difference >= 1 |subtrahend, difference| (subtrahend = minuend - difference)
		1. remove minuend from n
		2. newSet = [n..., subrahend, difference]
		3. concat(possible sets, combos(newSet)...)
4. return possible sets


n=[5]
[n]=[[]]


n=[1,1,1,1,1]
[n]=[[1,1,1,1,1]]

n=[1,1,1,2]
[n]=[[1,1,1,2], [1,1,1,1,1]]

*/