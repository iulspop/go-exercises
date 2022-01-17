package coin_change

import (
	"go-exercises/assert"
	"testing"
)

func TestCoinChange(t *testing.T) {
	result := CoinChange([]int {186,419,83,408}, 6249)
	expected := 20
	assert.Equals(t, "CoinChange()", "", "", result, expected)
}
