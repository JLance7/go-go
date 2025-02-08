package test
// package main

import "fmt"
import "math"

// import {
// 	"math"
// }

func main(){
	fmt.Println("Hello World")
	fmt.Printf("Sum is %v, diff is %v", add(1, 2), sub(1, 2))
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	flag := false // var flag bool = false
	var i = 0
	for ; flag == false; {
		i += 1
		if i == 50 {
			flag = true
		}
	}
	fmt.Println("done")

	// while loop
	for flag == true {
		// fmt.Println("in loop")
		i += 1
		if i == 100 {
			flag = false
		}
	}
	fmt.Printf("flag is: %v\n", flag)
	fmt.Println(sqrt(10))
	fmt.Println(pow(3, 2, 8))
}



func add(x int, y int) int {
	var sum int = 0
	sum = x + y
	return sum
}

func sub(a int, b int) int {
	sum := 0
	sum = a - b
	return sum
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("v is less than lim")
		return v
	}
	fmt.Println("v is greater than lim")
	return lim
}