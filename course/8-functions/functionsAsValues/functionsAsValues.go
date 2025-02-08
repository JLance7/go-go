package functionsAsValues

import (
	"fmt"
)

type transfromFn func(int) int

func main(){
	numbers := []int{1, 2, 3, 4}
	// doubleNumbersByReference(&numbers)
	// numbers = doubleNumbersByValue(numbers)
	numbers = transformNumbers(&numbers, double)
	fmt.Println(numbers)

	numbers = transformNumbers(&numbers, triple)
	fmt.Println(numbers)
}

// func doubleNumbersByReference(numbers *[]int){
// 	for index, val := range *numbers {
// 		(*numbers)[index] = double(val)
// 	}
// }

// func doubleNumbersByValue(numbers []int) []int {
// 	newNumbers := []int{}
// 	for _, val := range numbers {
// 		newNumbers = append(newNumbers, transformNumbers(numbers, double())...)
// 	}
// 	return newNumbers
// }

func transformNumbers(numbers *[]int, transform transfromFn) []int {
	newNumbers := []int{}
	for i, val := range *numbers {
		// newNumbers = append(newNumbers, transform(val))

		multiplier := getTransformerFunction(i+1)
		newNumbers = append(newNumbers, multiplier(val))
	}
	return newNumbers
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func getTransformerFunction(num int) transfromFn {
	if num % 2 == 0 {
		return double
	}
	return triple
}