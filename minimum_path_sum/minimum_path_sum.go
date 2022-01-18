package minimum_path_sum

func MinPathSum(grid [][]int) int {
	dp := map[[2]int]int {{0, 0}: grid[0][0]}

	rows := len(grid) - 1
	columns := len(grid[0]) -1

	for row := 0; row <= rows; row++ {
		for column := 0; column <= columns; column++ {
			if row == 0 && column == 0 { continue }
			if row == 0 {
				dp[[2]int {row, column}] = grid[row][column] + dp[[2]int {row, column - 1}] 
			} else if column == 0 {
				dp[[2]int {row, column}] = grid[row][column] + dp[[2]int {row - 1, column}] 
			} else {
				dp[[2]int {row, column}] = grid[row][column] + min(dp[[2]int {row - 1, column}],  dp[[2]int {row, column - 1}])
			}
		}
	}

	return dp[[2]int {rows, columns}]
}

func min(num1 int, num2 int) int {
	if num1 <= num2 { return num1 }
	return num2
}


/*


dp[0][0] = grid[0][0]

iterate rows
  iterate columns
    if 0,0 continue
    dp[row][column] = grid[row][colum] + min(dp[row - 1][column], dp[row][column + 1])
    

for any cell, the minimum is itself plus minimum of path above and to left


return dp[rows][columns]

*/