package main

import (
	"fmt"
	"os"

	multipleofthreeandfive "github.com/GaijinZ/task-2/multiple-of-three-and-five"
	negativenumber "github.com/GaijinZ/task-2/negative-number"
	reverseDigits "github.com/GaijinZ/task-2/reverse-digits"
)

func GetArr(message string, arr []int) []int {

	if message != "" {
		fmt.Println(message)
	}

	var number int
	n, err := fmt.Scanln(&number)

	if err != nil || n == 0 {
		return arr
	}

	arr = append(arr, number)
	return GetArr("", arr)
}

func GetInt(message string) int {

	fmt.Println(message)

	var number int
	_, err := fmt.Scanln(&number)

	if err != nil {
		fmt.Println("Error")
		return 0
	}

	return number
}

func main() {
	// fmt.Println("Pick a function to run: \n 1: Reverse Digits \n 2: Numbers that are multiple of 3 and not of 5 \n 3: Negative Number")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "1":
			for {
				numbers := GetInt("Enter numbers to reverse them")
				fmt.Println(reverseDigits.ReverseDigits(numbers))
			}
		case "2":
			for {
				arr := []int{}
				arr = GetArr("Input values to array", arr)
				fmt.Println(multipleofthreeandfive.MultipleOfThreeaAndFive(arr))
			}
		case "3":
			for {
				arr := []int{}
				arr = GetArr("Input values to array", arr)
				fmt.Println(negativenumber.NegativeNumber(arr))
			}
		}
	}

	// fmt.Println("Numbers that are multiple by 3 and not by 5: \n", multipleofthreeandfive.MultipleOfThreeaAndFive(nums))
	// fmt.Println("Numbers backwards: \n", reverseDigits.ReverseDigits(1))
	// fmt.Println("Numbers without negative number: \n", negativenumber.NegativeNumber(nums))
}
