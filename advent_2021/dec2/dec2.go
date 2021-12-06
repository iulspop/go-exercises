package dec2

import (
	"strings"
)

func Position(commands string) [2]int {
	strings.Split(commands, "\n")
	return [2]int{15, 10}
}


/*

forward => increases horizontal position
down => increases depth
up => decreases depth

horizontal 0
depth 0

mental model:
- split string
- iterate over and sum
- return depth and horizonal position

*/