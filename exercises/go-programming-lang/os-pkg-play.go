package turtle

import (
	"fmt"
	"os"
	"strings"
)

func turtle() {
	fmt.Print(strings.Join(os.Args, ","))
}