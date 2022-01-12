package minimum_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	for length := 1; length <= len(nums); length++ {
		for start := 0; start <= (len(nums) - length); start++ {
			window := nums[start:(start + length)]
			sum := sumInt(window)
			if sum >= target { return length }
		}
	}
	return 0
}

func sumInt(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}


/*

1. iterate from smallest slice (len = 1) to largest slice (len = nums.length) |length|
 1. iterate from 0 to <= (nums.length - array.length) |start index|
		1. window = nums[start:start + length]
		2. sum window
		3. if sum >= target, if true return length
2. return 0


Input: nums , target    
Output: minimum length

Explicit Requirements:
- Contiguous subarray, index to index + x, without gaps
- smallest subarray greater or equal to target
- if no subarrary equal or greater to target, return 0
- only positve int inputs

Implicit Requirements:
- assume no 0 target
- assume nums slice isn't empty

Test Cases
two elements
- 7, [2,3,1,2,4,3] => [4,3] => 2

only element enough
- 4, [1, 4, 4] => [4] => 1

not found
- 11, [1, 1, 1, 1, 1, 1, 1] => ? => 0

Mental model:
[]
[ , ]
[ , , ]
until you find first sum greater or equal, or return

iterate from smallest slice (len = 1) to largest slice (len = nums.length)
iterate from first index to last index that fits
	check sum >= target, if true return length
	
	

Algo
1. iterate from smallest slice (len = 1) to largest slice (len = nums.length) |length|
 1. iterate from 0 to <= (nums.length - array.length) |start index|
		1. window = nums[start:start + length]
		2. sum window
		3. if sum >= target, if true return length
2. return 0


*/