package intersect

func Intersect(slice1 []int, slice2 []int) []int {
	numbers := make(map[int]bool)

  for i := 0; i < len(slice1); i++ {
		numbers[slice1[i]] = true
	}


	var intersect []int
	for i := 0; i < len(slice2); i++ {
		number := slice2[i]

		if numbers[number] { intersect = append(intersect, number) }
	}

	return intersect
}
