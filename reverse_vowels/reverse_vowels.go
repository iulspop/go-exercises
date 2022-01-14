package reverse_vowels

var vowels = map[rune]bool {
  'a': true,
  'e': true,
  'i': true,
  'o': true,
  'u': true,
	'A': true,
  'E': true,
  'I': true,
  'O': true,
  'U': true,
}

func ReverseVowels(s string) string {
  runes := []rune(s)
  start, end := 0, len(runes) - 1

  for start < end {
    if vowels[runes[start]] && vowels[runes[end]] {
      runes[start], runes[end] = runes[end], runes[start]
      start++
			end--
    }
    if !vowels[runes[start]] { start++ }
    if !vowels[runes[end]]   { end-- }
  }

  return string(runes)
}

/*

input: string
ouput: new string with vowels in reverse order

explicit req:
- 'a', 'e', 'i', 'o', and 'u' are vowels
- string at least one length
- string is ASCII chars!!!

implcit req:
- inplace or build new string => assume I can do "inplace"
- ASCII chars means both lowercase and uppercase

test cases:
h0lle => holle
  ! 
  !

leotcedes => leotcede
    s
   e
   
e => e
s
e


mental model:
two pointers: start, end

when move start?
 if not vowel, increase
when move end?
 if not vowel, decrease

on what condition do I switch letters?
  if both vowels swap (swap indexes)
  increase start, decrease end

on what condition to I end?
  start >= end
  
algo:
- create array of runes
- init start, end = 0, len() - 1
- for start < end
   - if both vowels
     - swap
     - increase start, decrease end
   - if start not vowel
     - increase
   - if end not vowel
     - decrease
- concat array of runes into string and return

""

*/