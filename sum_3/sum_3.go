package sum_3

import (
	"fmt"
	"go-exercises/insertion_sort"
)

func Sum3(nums []int) [][]int {
	indexes := indexNums(nums)

	triplets := [][]int {}
	tripletsIndex := map[string]bool {}
	for negative := range indexes.negatives {
		for positive := range indexes.positives {
			difference := -1 * (negative + positive)

			if difference > 0 {
				count, ok := indexes.positives[difference]

				if ok {
					if difference == positive {
						if count >= 2 {
							triplets = insertIfNotIndexed(tripletsIndex, triplets, []int {negative, difference, positive})
						}
					} else {
						triplets = insertIfNotIndexed(tripletsIndex, triplets, []int {negative, difference, positive})
					}
				}
			}

			if difference < 0 {
				count, ok := indexes.negatives[difference]

				if ok {
					if difference == negative {
						if count >= 2 {
							triplets = insertIfNotIndexed(tripletsIndex, triplets, []int {negative, difference, positive})
						}
					} else {
						triplets = insertIfNotIndexed(tripletsIndex, triplets, []int {negative, difference, positive})
					}
				}
			}

			if difference == 0 {
				_, ok := indexes.zeros[difference]

				if ok {
					triplets = insertIfNotIndexed(tripletsIndex, triplets, []int {negative, difference, positive})
				}
			}
		}
	}

	if indexes.zeros[0] >= 3 { triplets = append(triplets, []int {0, 0, 0}) }

  return triplets
}

type numsIndexes struct {
	positives map[int]int
	negatives map[int]int
	zeros    map[int]int
}

func indexNums(nums []int) numsIndexes {
	positives := map[int]int{}
	negatives := map[int]int{}
	zeros := map[int]int{}

	for _, num := range nums {
			if num > 0 {
				positives[num] += 1
			}

			if num < 0 {
				negatives[num] += 1
			}

			if num == 0 {
				zeros[num] += 1
			}
	}

	return numsIndexes{positives, negatives, zeros}
}

func insertIfNotIndexed(index map[string]bool, triplets [][]int, triplet []int) [][]int {
	key := fmt.Sprint(insertion_sort.InsertionSort(triplet))
	_, ok := index[key]

	if !ok {
		index[key] = true
		return append(triplets, triplet)
	} else {
		return triplets
	}
}

/*

Find all the triplets
1. init triplets
2. Iterate over keys (negative num) in negative map
   1. Iterate over keys (postive num) in positive map
      1. negative + positive = difference
      2. if difference positive, then check if exist in positive hash
         1. it exists if is found and if it is the same positive number, then the count must be greater than 2
      3. if difference negative, then ...
      4. if difference 0, then ...
      5. if difference found then add combo [positive, difference, negative]
3. if zeros.count >= 3, add {0, 0, 0}
4. return triplets
*/


/*


input: []int
output: [][]int all triplets

explicit requirements:
- i != j != k (distinct elements in slice)
- sum of elements == 0
- set of unique combination (no duplicates)

implicit requirements:
- assume if less than three elements return empty slice
- combination can be three distict 0
- order matters?

edge cases
[] => []
[0] => [] (bcs no three elements)



[-1,0,1,2,-1,-4]

{
  -1: 2
   0: 1
   1: 1
   2: 1
   -4: 1
}

for every negative number, need a corresponding postive number
 
 
 input: {-int: count}(only negative) {int: count}(positive including 0)
 output: all the combinations that sum to 0
 
 -1
   1
   0
   
   -1 + x = 0 
    x = 0 - -1
    x = 1
    
    -1 + 54 - 53
   
 -2
   2
   0

   1
   1
   
-3
  3
  0
  
  2
  1


-1

  -50
  
        1
-1  2
        1
  
 
iterate over negative integers in map
  iterate over positive integers in map (not zeros)
    only one possible positive/negative match
     if exists add combo

check if three distinct 0
  
1 2 3

-1 -2 -3
 
middle any other
  
  
Algo:

Index num counts in two maps
1. init negative nums map
2. init positive nums map
3. init zeros count map
3. iterate over nums
    1. if num negative
       1. if key doesn't exist, then init to 1
       2. if key does exit, increment by 1
    2. if num positive
       1. if key doesn't exist, then init to 1
       2. if key does exit, increment by 1
    3`. if num positive
       1. if key doesn't exist, then init to 1
       2. if key does exit, increment by 1
4. return maps

Find all the triplets
1. init triplets
2. Iterate over keys (negative num) in negative map
   1. Iterate over keys (postive num) in positive map
      1. negative + positive = difference
      2. if difference positive, then check if exist in positive hash
         1. it exists if is found and if it is the same positive number, then the count must be greater than 2
      3. if difference negative, then ...
      4. if difference 0, then ...
      5. if difference found then add combo [positive, difference, negative]
3. if zeros.count >= 3, add {0, 0, 0}
4. return triplets

*/
    