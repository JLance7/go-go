package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// anonymous function
	transformed := transformNumbers(&numbers, func (num int) int {
		return num * 4
	})

	double := createTransformer(2)
	triple := createTransformer(3)
	transformed2 := transformNumbers(&numbers, double)
	transformed3 := transformNumbers(&numbers, triple)
	fmt.Println(transformed2)
	fmt.Println(transformed3)


	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// produces functions
func createTransformer(factor int) func(int) int {
	return func(num int) int{
		return num * factor
	}
}