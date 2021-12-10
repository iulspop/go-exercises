package binary_search

import "math"

func BinarySearch(numbers []int, searchValue int) int  {
	middleIndex := len(numbers) / 2

	for i := 0; i <= log2Int(len(numbers)); i++ {
		value := numbers[middleIndex]
		if value == searchValue { return middleIndex }
		if value < searchValue  { middleIndex = len(numbers) - middleIndex / 2  }
		if value > searchValue  { middleIndex = middleIndex / 2 }
	}

	return -1
}

func log2Int(num int) int {
	return int(math.Floor(math.Log2(float64(num))))
}