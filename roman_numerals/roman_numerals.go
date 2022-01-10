package roman_numerals

import (
	"go-exercises/slice_stack"
)

func RomanToInt(chars string) string {

}

/*
Input: string
Output: int

Explicit Req:
- 7 symbols
- Add largest left to right
- If a smaller symbol precedes larger, means it substracts from larger
    - 6 cases where that happens
- max length 15, at least 1
- only contains valid symbols
- guaranteed valid roman numeral

Test Cases:
  romanToInt("III") // === 3
  
  romanToInt("LVIII") // === 58
  
  Substraction
  romanToInt("MCMXCIV") // === 1994
  M 1000
  C
  
  
  C, M
  X, C
  I, V

Mental Model:

FIRST TRY:
init total
init previous symbol
iterate from left to right |current symbol|
  if current I X or C then next

  if I is previous and is current V or X
    add remainder of substraction
  if X is previous and is current L or C
    add sub
  if C is previous and is current D or M
    add sub
  if previous is I X or C
    add current & previous
  else
    add current
return total

SIMPLIFIED:
first map symbols to slice of integers
iterate of slice
  if next is greater than current add substraction and skip iteration
  else add to total

*/