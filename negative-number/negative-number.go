package negativenumber

func NegativeNumber(arr []int) []int {
	results := []int{}

	for index, value := range arr {
		if value < 0 {
			break
		}

		results = append(results, value*arr[index+1])
	}

	return results
}
