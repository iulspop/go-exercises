package combinations

func Combine(n int, k int) [][]int {
  list := makeList(n)
  canditate := []int {}
  result := [][]int {}

  backtrack(k, list, canditate, &result)
  return result
}

func backtrack(k int, list []int, canditate[]int, result *[][]int) {
  if len(canditate) == k {
    canditateCopy := make([]int, len(canditate))
    copy(canditateCopy, canditate)
    *result = append(*result, canditateCopy)
    return
  }

  for i, num := range list {
    canditate = append(canditate, num)
    backtrack(k, list[i+1:], canditate, result)
    canditate = canditate[:len(canditate)-1]
  }
}

func makeList(n int) []int {
  list := []int {}
  for i := 1; i <= n; i++ {
    list = append(list, i)
  }
  return list
}


/*

input: n k
output: possible comvinations

explicit req:
- a combination has k nums
- 1 to n is the range

implicit req:
- you can use numbers in n range only once
- 


list = 1 to n
candidate = []
result = [][]int





*/
