package reversedigits

func ReverseDigits(number int) int {
	res := 0

	for number > 0 {
		remainder := number % 10
		res = (res * 10) + remainder
		number /= 10
	}

	return res
}
