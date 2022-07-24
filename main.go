package main

import (
	"fmt"
	"os"
	"strconv"

	multipleofthreeandfive "github.com/GaijinZ/task-2/multiple-of-three-and-five"
	negativenumber "github.com/GaijinZ/task-2/negative-number"
	reverseDigits "github.com/GaijinZ/task-2/reverse-digits"
)

func main() {
	commandLineArguments := os.Args[1:]
	var err error
	nums := make([]int, len(commandLineArguments))

	for i := 0; i < len(commandLineArguments); i++ {
		if nums[i], err = strconv.Atoi(commandLineArguments[i]); err != nil {
			panic(err)
		}
	}

	fmt.Println("Numbers that are multiple by 3 and not by 5: \n", multipleofthreeandfive.MultipleOfThreeaAndFive(nums))
	fmt.Println("Numbers backwards: \n", reverseDigits.ReverseDigits(-10))
	fmt.Println("Numbers without negative number: \n", negativenumber.NegativeNumber(nums))
}
