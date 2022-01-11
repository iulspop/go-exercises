package is_subsequence

func IsSubsequence(s string, t string) bool {
	anchor := 0
	trunes := []rune(t)

	for _, srune := range s {
		for i := anchor; i < len(trunes); i++ {
			trune := trunes[i]
			if srune == trune { anchor = i + 1; break }
			if i == len(trunes) - 1 { return false }
		}
	}

	return true
}

/*

Input: s, t
Output: boolean

Explicit requirements:
- is s a subsequence of t
- is s a subsequence if we remove some characters (or none) from `t` can we form `s`, without distrubing the relative positions of the chars in `t`. (so characters can be deleted but order stays the same).

Mental Model:

 
anchor 1

 "abc" <= s
  !
  
 "ahbgdc" <= t
  
iterate over s chars
  check if char occurs in t after anchor
    move anchor

Data Structures:
just pointers

Algorightm:

1. init anchor = 0
2. iterate over s runes
   1. iterate over t runes from anchor
      1. check if t char == s char
         1. move anchor to postion next t char
         2. break
      2. if last index then return false
3. return true


*/