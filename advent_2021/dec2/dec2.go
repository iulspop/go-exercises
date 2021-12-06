package dec2

import (
	"regexp"
	"strconv"
	"strings"
)

func Position(commandsString string) [2]int {
	commands := strings.Split(commandsString, "\n")
	horizontalPosition := 0
	depth := 0

	forward, _ := regexp.Compile("forward")
	up, _ := regexp.Compile("up")
	down, _ := regexp.Compile("down")
	num, _ := regexp.Compile("[0-9]")

	for _, command := range commands {
		distance, _ := strconv.Atoi(num.FindString(command))

		if forward.MatchString(command) {
			horizontalPosition += distance
		} else if up.MatchString(command) {
			depth -= distance
		} else if down.MatchString(command) {
			depth += distance
		}
	}

	return [2]int{horizontalPosition, depth}
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