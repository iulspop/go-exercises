package roman_numerals

var numerals = map[rune]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(chars string) int {
	runes := []rune(chars)

	sum := 0
	for i := 0; i < len(runes); i++ {
		current := numerals[runes[i]]
		if i == len(runes) -1 { sum += current; break }
		next := numerals[runes[i + 1]]

		if next > current {
			sum += (next - current); i++
		} else {
			sum += current
		}
	}

	return sum
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