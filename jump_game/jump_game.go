package jump_game

func CanJump(nums []int) bool {
  if len(nums) == 1 { return true }

  var helper func(int) bool
  helper = func(index int) bool {
    num := nums[index]
    for jump := num; jump >= 1; jump-- {
      if (index + jump) >= len(nums) - 1 { return true }
      if helper(index + jump) { return true }
    }
    return false
  }

  return helper(0)
}

/*
algo:
- define helper as closure including nums
- call helper

helper(index)
  - num := nums[index]
  - iterate from num to 1 inclusive |jump|
    - if (index + jump) >= len(nums) - 1 { return true }
    - if helper(index + jump) { return true }
  - return false



input: nums
output: 

explicit requirements:
- start at first index (0)
- each element is **MAX** jump length from that position
- return true if you can reach the last index (land on last but not over?)
- array contains at least on element
- numbers 0 or positive

implicit requirements:
- there is never a case where you must go over and not land on last index, since if you can jump over, you can jump one less.

test cases:
[2,3,1,1,4]


       2
    2      1
  

mental model:
start with first element
branch paths from num to 1
  in each path repeat branching
  if reach last or over, return true
  
if try all paths and nothing, return false


algo:
- define helper as closure including nu,s
- call helper

helper(index)
  - num := nums[index]
  - iterate from num to 1 inclusive |jump|
    - if (index + jump) >= len(nums) - 1 { return true }
    - if helper(index + jump) { return true }
  - return false







*/