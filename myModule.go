package myModule

import "fmt"

func Demo(x int) int {
	return x * x
}

func Demo2(x int) int {
	return x * x * x
}

func Version() {
	fmt.Println("Version 2.0.1")
}
