package main

import (
	"fmt"
	"task-golang/utils" // Import the utils package
)

func main() {

	n := 1

	fmt.Println(utils.MagicSum(n))
	fmt.Println(utils.MagicPow(n))
	fmt.Println(utils.MagicOdd(n))
	fmt.Println(utils.MagicGrade(n))
	fmt.Println(utils.MagicName(n))

	// MagicChange
	fmt.Println("Before MagicChange:", n)
	utils.MagicChange(&n)
	fmt.Println("After MagicChange:", n)

	// MagicTria
	fmt.Println("MagicTria:", utils.MagicTria(n))

	// Initialize MagicNumber struct and use Multiply method
	magicNumber := utils.MagicNumber{Number: n}
	fmt.Println("Before Multiply:", magicNumber.Number)
	magicNumber.Multiply(n)
	fmt.Println("After Multiply:", magicNumber.Number)
}
