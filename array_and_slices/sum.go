package main

// import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfnumbers := len(numbersToSum)
	// fmt.Println(lengthOfnumbers)
	sums := make([]int, lengthOfnumbers)
	// fmt.Println(sums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
		// fmt.Println(i, numbers, sums[i])
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func main() {
	SumAll([]int{1, 2}, []int{0, 9})
}
