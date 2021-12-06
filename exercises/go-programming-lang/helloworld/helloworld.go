package helloworld

import (
	"fmt"
	"os"
	"strings"
)

func Hello() {
	fmt.Print(fmt.Errorf("got %v, wanted %v", 4, 5).Error())
	fmt.Print(strings.Join(os.Args, ","))
}
