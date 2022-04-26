package main

import (
	"fmt"
)

func main() {
	userInput := 0
	var newResult float64

	for {
		fmt.Println("Input a number")
		fmt.Scanln(&userInput)
		if userInput%2 == 1 {
			fmt.Println("This number is odd")

		} else {
			fmt.Println("This number is even")
		}
		newResult = float64(userInput) / 10
		if newResult > 0.9 || newResult < -0.9 {
			fmt.Println("Number has more than 1 digit!", newResult)
		} else {
			fmt.Println("Number is only 1 digit!", newResult)
		}
	}
}
