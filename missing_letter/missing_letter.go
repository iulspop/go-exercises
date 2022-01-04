package missing_letter

func MissingLetter(phrase string) string {
	alphabetIndex:= map[string]bool{
		"a": false,
		"b": false,
		"c": false,
		"d": false,
		"e": false,
		"f": false,
		"g": false,
		"h": false,
		"k": false,
		"l": false,
		}


	for _, letter := range phrase {
			alphabetIndex[string(letter)] = true
	}

	for letter, seen := range alphabetIndex {
		if !seen { return letter }
	}

	return ""
}
