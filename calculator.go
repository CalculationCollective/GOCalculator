package main

import "fmt"

func main() {
	var first float64
	var f int
	var second float64

	fmt.Print("Please enter first number: ")
	fmt.Scan(&first)

	fmt.Println("Choose between:\n1)Addition\n2)Subtraction\n3)Multiplication\n4)Division")
	fmt.Scan(&f)

	fmt.Print("Please enter second number: ")
	fmt.Scan(&second)

	switch f {
	case 1:
		fmt.Println(first + second)
	case 2:
		fmt.Println(first - second)
	case 3:
		fmt.Println(first * second)
	case 4:
		fmt.Println(first / second)
	default:
		fmt.Println("No valid function chosen")
	}
}
