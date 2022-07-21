package multipleofthreeandfive

func MultipleOfThreeaAndFive(arr []int) []int {
	finalNumbers := []int{}

	for _, element := range arr {
		if element%3 == 0 && element%5 != 0 {
			finalNumbers = append(finalNumbers, element)
		}
	}

	return finalNumbers
}
