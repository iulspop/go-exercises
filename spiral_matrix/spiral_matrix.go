package spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	columns := len(matrix[0])

	top := 0
	right := columns - 1
	bottom := rows - 1
	left := 0

	result := []int {}
	for {
		if (left > right) { break }
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		if (top > bottom) { break }
		for i := top; i <= bottom; i++ {
		result = append(result, matrix[i][right])
		}
		right--

		if (left > right) { break }
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		if (top > bottom) { break }
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
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