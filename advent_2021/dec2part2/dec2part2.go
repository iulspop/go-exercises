package dec2part2

import (
	"regexp"
	"strconv"
	"strings"
)

func Position(commandsString string) [2]int {
	commands := strings.Split(commandsString, "\n")
	horizontalPosition := 0
	depth := 0
	aim := 0

	forward, _ := regexp.Compile("forward")
	up, _ := regexp.Compile("up")
	down, _ := regexp.Compile("down")
	num, _ := regexp.Compile("[0-9]")

	for _, command := range commands {
		units, _ := strconv.Atoi(num.FindString(command))

		if forward.MatchString(command) {
			horizontalPosition += units
			depth += (units * aim)
		} else if up.MatchString(command) {
			aim -= units
		} else if down.MatchString(command) {
			aim += units
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