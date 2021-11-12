package myModule

import "fmt"

func Demo(x int) int {
	return x * x
}

func Demo2(x int) int {
	return x * x * x
}

func Sum(x int, y int) int {
	return x + y
}

func Version() {
	fmt.Println("Version 2.1.0")
}
