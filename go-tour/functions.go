package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

// shorthand types
func shortHandTypes(x, y int) int {
	return x + y
}

// return types
func returnTypes(x, y string) (string, string) {
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func nakedReturn() {
	fmt.Println(split(17)) // 7 10
}
