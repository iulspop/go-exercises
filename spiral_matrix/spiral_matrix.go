package spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	columns := len(matrix[0])

	top := 0
	right := columns - 1
	bottom := rows - 1
	left := 0

	total := rows * columns
	result := []int {}
	for {
		to(left, right, func(current int)  {
			result = append(result, matrix[top][current])
		})
		top++
		if (len(result) >= total) { break }

		to(top, bottom, func(current int)  {
			result = append(result, matrix[current][right])
		})
		right--
		if (len(result) >= total) { break }

		downto(right, left, func(current int)  {
			result = append(result, matrix[bottom][current])
		})
		bottom--
		if (len(result) >= total) { break }

		downto(bottom, top, func(current int)  {
			result = append(result, matrix[current][left])
		})
		left++
		if (len(result) >= total) { break }
	}

	return result
}

func to(from int, to int, callback func(int))  {
	if from > to { return }
	for i := from; i <= to; i++ {
		callback(i)
	}
}

func downto(from int, to int, callback func(int))  {
	if from < to { return }
	for i := from; i >= to; i-- {
		callback(i)
	}
}

/*
Input: matrix (n * m!)
Output: single slice

Explicit requirements: m * n matrix!

Mental Model:
top
right
bottom
left

loop
	left to right
	top to bottom
	right to left
	bottom to top

add elements to array as I go

Data Structures:
init slice []

Algo:

rows = length of top array
columns = length of subarrays

top: 0
bottom: rows - 1
left: 0
right: columns - 1

loop until result.length = rows * columns
	left to right |current|
		grab [top, current]
	top++

	top to bottom |current|
		grab [current, right]
	right--

	right to left |current|
		grab [bottom, current]
	bottom--

	bottom to top |current|
	  grab [current, left]
	left++

1, 2
4, 3

1, 2, 3
      4
      5



*/