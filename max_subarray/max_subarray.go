package max_subarray

import "math"

func MaxSubArray(nums []int) int {
  if len(nums) == 1 { return nums[0] }

  left   := MaxSubArray(nums[:len(nums)/2])
  right  := MaxSubArray(nums[len(nums)/2:])
  center := maxCenter(nums)

  return maxNum(left, center, right)
}

func maxCenter(nums []int) int {
  maxLeft, maxRight := -math.MaxInt32, -math.MaxInt32

  sum := 0
  for i := len(nums)/2 - 1; i >= 0; i-- {
    sum += nums[i]
    if sum > maxLeft { maxLeft = sum }
  }

  sum = 0
  for i := len(nums)/2; i < len(nums); i++ {
    sum += nums[i]
    if sum > maxRight { maxRight = sum }
  }

  return maxLeft + maxRight
}

func maxNum(num1 int, num2 int, num3 int) int {
  if num1 >= num2 && num1 >= num3 { return num1 }
  if num2 >= num1 && num2 >= num3 { return num2 }
  return num3
}


/*
maxCenter(nums)
  - init maxLeft, maxRight = -infinity
  
  - init sum
  - for middle to start   middle = len(nums)/2 - 1    start = 0
    - add to sum
    - check if greater than maxLeft, then increment
      
  - reset sum
  - for middle to end     middle = len(nums)/2        end = len -1
     - add to sum
     - check if greater than maxRight, then increment  
     
  - return maxLeft + maxRigth





input: nums
output: max sum



if len == 1 return nums[0]

return maxNum(left, center, right)


maxSub(nums)

left = :len(nums) / 2 

right = len(nums) / 2:


[0,1, 2,3]
            len/2 = 2 -1      1 to 0
            len/2 = 2     2 to 3
    
[0,1, 2, 3,4]
               2 to 0
               3 to 4


maxSub(nums)
  - if len == 1 return nums[0]
  
  - left  = maxSub(nums[:len(nums)/2])
  - right = maxSub(nums[len(nums)/2:])
  - center = maxCenter(nums)

  - return maxNum(left, center, right)


maxCenter(nums)
  - init maxLeft, maxRight = -infinity
  
  - init sum
  - for middle to start   middle = len(nums)/2 - 1    start = 0
    - add to sum
    - check if greater than maxLeft, then increment
      
  - reset sum
  - for middle to end     middle = len(nums)/2        end = len -1
     - add to sum
     - check if greater than maxRight, then increment  
     
  - return maxLeft + maxRigth

maxNum(num1, num2, num3)
  if num1 > num2 && num1 > num3 return num1 
  if num2 > num1 && num2 > num3 return num2
  return num3




*/