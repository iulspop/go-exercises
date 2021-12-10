package bubble_sort

func BubbleSort(numbers []int) []int  {
	unsortedUntilIndex := len(numbers) - 1
	swap := true

	for swap {
		swap = false
		for i := 0; i < unsortedUntilIndex; i++ {
			first  := numbers[i]
			second := numbers[i + 1]

			if first > second { numbers[i] = second; numbers[i + 1] = first; swap = true }
		}
		unsortedUntilIndex--
	}

	return numbers
}
