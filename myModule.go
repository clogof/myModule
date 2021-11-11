package myModule

import "fmt"

func Demo(x int) int {
	return x * x
}

func Demo2(x int) int {
	return x * x * x
}

func Version() {
	fmt.Println("actual str")
	fmt.Println("Version 1.1.0")
}
