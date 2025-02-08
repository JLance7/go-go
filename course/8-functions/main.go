package main

import "fmt"

func main(){
	// fmt.Println(factorial(5))
	fmt.Println(sumNumbers(0, 1, 2, 3))

	numbersInSlice := []int{1, 2, 3}
	fmt.Println(sumNumbers(0, numbersInSlice...))
}

func factorial(num int) int{
	if num <= 1 {
		return 1
	} 
	return num * factorial(num - 1)
}

// variadic function
func sumNumbers(startingVal int, numbers ...int) int {
	sum := startingVal
	for _, val := range numbers {
		sum += val
	}
	return sum
}