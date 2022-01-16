package is_palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(str string) bool {
  str = strings.Join(regexp.MustCompile("([[:alpha:]]|\\d)*").FindAllString(strings.ToLower(str), -1),"")
  return checkLetterEqual(str, 1)
}

func checkLetterEqual(s string, n int) bool {
  if n == (len(s) / 2) + 1 { return true }
  return s[n - 1] == s[len(s) - n] && checkLetterEqual(s, n + 1)
}

/*

input string
output bool

explicit req:
- read the same forward and backward even if remove all non-alphanumeric charcters and spaces and all are lowercase
- [^a-bA-B0-9]

test cases:
"amanaplana c analpanama"

mental model:

recursively check from sides to center

if n = len / 2
if check if two letters same && check(string, n + 1)  {
  return true
} else {
 false
}

algo:

isPalindrome
  - clean string
  - return checkLettersSame(string, 1)
  
checkLettersSame
  string, n
  
  - if n = (len / 2) + 1 return true
  - if string[n - 1] == string[len - n] && checkLettersSame(string, n + 1)
    - true
    else
    - false


"abcba"
 1
  2 
   3
   
"abda"
 1
  2
   3
*/