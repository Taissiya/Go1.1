package main

import (
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

func min(a, b int) int {
	return a - b
}

func main() {
	var a int
	var b, c int
	fmt.Println("1:add or 2:mines ?")
	fmt.Fscan(os.Stdin, &a)
	switch a {
	case 1:
		fmt.Println("Enter 1 number")
		fmt.Fscan(os.Stdin, &b)
		fmt.Println("Enter 2 number")
		fmt.Fscan(os.Stdin, &c)
		fmt.Println(add(b, c))
	case 2:
		fmt.Println("Enter 1 number")
		fmt.Fscan(os.Stdin, &b)
		fmt.Println("Enter 2 number")
		fmt.Fscan(os.Stdin, &c)
		fmt.Println(min(b, c))

	default:
		fmt.Println("Panic")

	}
}
