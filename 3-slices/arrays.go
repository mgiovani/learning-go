package main

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(numbersToSum ...[]int) []int {
	lenghtOfNumbers := len(numbersToSum)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = sum(numbers)
	}

	return sums
}
