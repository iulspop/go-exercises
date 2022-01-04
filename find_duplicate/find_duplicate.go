package find_duplicate

func FindDuplicate(letters []string) string {
	lettersIndex := make(map[string]bool)

	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		if lettersIndex[letter] { return letter }
		lettersIndex[letter] = true
	}

	return ""
}
