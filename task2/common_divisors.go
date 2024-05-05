package findCommonDivisors



func findCommonDivisors(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	minNumber := numbers[0]
	for _, num := range numbers {
		if num < minNumber {
			minNumber = num
		}
	}

	divNum := make(map[int]int)

	for div := 2; div <= minNumber; div++ {
		isDiv := true
		for _, num := range numbers {
			if num % div != 0 {
				isDiv = false
				break
			}
		}
		if isDiv {
			divNum[div] = 1
		}
	}

	var commonDiv []int
	for div := range divNum {
		commonDiv = append(commonDiv, div)
	}

	return commonDiv
}
