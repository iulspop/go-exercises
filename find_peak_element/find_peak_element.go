package find_peak_element

func FindPeakElement(nums []int) int {
  left, right :=  0, len(nums) - 1

  if len(nums) == 1 { return 0 }
  if nums[left] > nums[left + 1]   { return left }
  if nums[right] > nums[right - 1] { return right }

  left += 1
  right -= 1

  for left <= right {
    mid := (left + right) / 2
    if isPeak(nums, mid) { return mid }
    if leftOrRight(nums, mid) {
      right = mid - 1
    } else {
      left = mid + 1
    }
  }

  return -1
}

func isPeak(nums []int, mid int) bool {
  if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1] { 
    return true
  } else { 
    return false
  }
}

func leftOrRight(nums []int, mid int) bool {
  if nums[mid - 1] > nums[mid] {
    return true
  } else {
    return false
  }
}

/*
func findPeakElement(nums []int) int {
  if len(nums) == 1 { return 0 }
  left, right :=  0, len(nums) - 1
  
  for left <= right {
    mid := (left + right) / 2
    if isPeak(nums, mid) { return mid }
    if leftOrRight(nums, mid) {
      right = mid - 1
    } else {
      left = mid + 1
    }
  }
  
  return -1
}

func isPeak(nums []int, mid int) bool {
  if mid == 0 { 
    if nums[mid] > nums[mid + 1] {
      return true
    } else {
      return false
    }
  }

  if mid == len(nums) - 1 {
    if nums[mid] > nums[mid - 1] {
      return true
    } else {
      return false
    }
  }

  if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1] { 
    return true
  } else { 
    return false
  }
}

func leftOrRight(nums []int, mid int) bool {
  if mid == 0  { return false }

  if nums[mid - 1] > nums[mid] {
    return true
  } else {
    return false
  }
}

*/


/*
Input: nums
Output: return index of any peak

Explicit Requirement:
- peak element = greater than left or right
- i = -1 = -infinity and i = len() = infinity
- nums contains at least one element
- nums can include negative int
- any given element is not equal to it's neigbours

Implicit Requirements:
- 

Test Cases:
{1,2,3,1}
   !  

2

{1,2,1,3,5,6,4}
1 or 5

{1, 2, 3, 4, 5}
             !
             L
             R

get mid
  check if is peak (greater than both neighours)
  move left right towards increasing value
  
  
Algorithm:
1. init left, right = 0, len() - 1
2. for left <= right
   1. init mid = (left + right) / 2
   2. if isPeak(nums, mid) { return mid }
   3. if leftOrRight(nums, mid)
      1. right = mid - 1
      else
      1. left = mid + 1
3. return -1

isPeak(nums, mid)
 if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1] { return true }
 else { return false }
 

leftOrRight(nums, mid)
  if nums[mid - 1] > nums[mid]
    left
  else
    right

*/