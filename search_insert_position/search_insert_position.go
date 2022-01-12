package search_insert_position

func searchInsert(nums []int, target int) int {
  left, right := 0, len(nums) - 1

  for left <= right {
    mid := (left + right) / 2

    if nums[mid] == target { return mid }

    if leftOrRight(nums[mid], target) {
      right = mid - 1
    } else {
      left = mid + 1
    }
  }

  return left
}

func leftOrRight(value int, target int) bool {
  if value > target {
		return true
	} else {
		return false
	}
}

/*


input: []int, target
output: index of target or insert index

Explicit constraints:
- input numbers contains at least one element
- input numbers are sorted ascending
- input numbers are distict values (no repeats, so one place to insert and no multiple targets to be found)

Implicit requirements:
- you insert at the position where it will go, what occupied it shifts forward
- numbers can be negative
- target can be negative
- if target is greater than greatest value, or less than smallest value, insert at the ends of array

Teat Cases:
[]int {1, 3, 5, 6}, 5
2

[]int {1, 3, 5, 6}, 2
1

[]int {1, 3, 5, 6}, 7
4

Solution Mental Model:

left: 0
right: len(nums) - 1

mid = (left + right) / 2

<- ->
check value against target ? left : right
  
until find target || left == right

Algo:
1. init left = 0
2. init right = len(nums) - 1
3. init mid
4. for left <= right
   1. mid = (left + right) / 2
   1. value = nums[mid]
   2. if value == target then break
   3. if leftOrRight(value)
      1. right = mid - 1
      else
      1. left = mid + 1
5. return mid (at this point mid either target or insert index)

leftOrRight(value)
  1. if value > target
     1. return true (left)
     else
     1. return false (right)

[1, 3, 5, 6], 0



*/