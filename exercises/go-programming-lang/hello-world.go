package main

import "fmt"

func main() {
	fmt.Print(fmt.Errorf("got %v, wanted %v", 4, 5).Error())
}
