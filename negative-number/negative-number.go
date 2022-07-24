package negativenumber

func NegativeNumber(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	positiveNumbers := []int{}

	var finalResult int

	for index, value := range arr {
		if value < 0 || len(arr)-1 == index || arr[index+1] < 0 {
			break
		}

		positiveNumbers = append(positiveNumbers, value*arr[index+1])
	}

	for _, value := range positiveNumbers {
		finalResult += value
	}

	return finalResult
}
