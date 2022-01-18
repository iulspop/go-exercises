package coin_change

import (
	"sort"
	// "fmt"
)

func CoinChange(coins []int, amount int) int {
  if amount == 0 { return 0 }

  sort.Ints(coins)

  var helper func(int, int) int
  helper = func(amount int, count int) int {
    for i := len(coins)-1; i >= 0; i-- {
      coin := coins[i]
      if coin > amount { continue }
      if coin == amount { return count + 1 }

			// fmt.Println("decrement: ", coin)
			// fmt.Println("new amount: ", amount - coin, "new count: ", count + 1)
      result := helper(amount - coin, count + 1)
      if result != -1 { return result }
    }
    return -1
  }

  return helper(amount, 0)
}

/*

input: coins, amount
output: fewest num of coins

explicit requirements:
- coins represent possible denominations
- have no limit to how much of any denomination can use
- return ***fewest num of cons needed***
- return -1 if any combination doesn't add to amount
- coins include at least one
- amount may be zero or greater
- coin denominations are 1 or greater

implicit requirements:
- all integers
- if amount is zero, then zero coins are needed
- coins sorted ascending?


test cases:
[1, 2, 5] 11
3

[1], 0
0

[2], 3
-1

mental model:
take amount
  iterate over denominations
    if denomination greater than amount, then continue
    if denomination == amount, return count + 1
    return helper(amount - denomination, count + 1)

algo:
coinChange
- if amount == 0, return 0
- init helper (with coins in closure)
- return helper(amount, 0)

helper(amount, count)
- iterate over coins in reverse order
  - if coin > amount, continue
  - if coin == amount, return count + 1

  - result = helper(amount - denomination, count + 1)
  - if result != -1 { return result }
- return -1
    
[1, 2, 5] 11
3


*/