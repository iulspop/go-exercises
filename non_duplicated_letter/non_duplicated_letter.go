package non_duplicated_letter

func NonDuplicatedLetter(phrase string) string {
	letterIndex:= make(map[string]int)

	for _, letter := range phrase {
			letterIndex[string(letter)] += 1
	}

	for letter, count := range letterIndex {
		if count < 2 { return letter }
	}

	return ""
}
