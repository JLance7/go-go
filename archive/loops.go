package main

import (
	"fmt"
	"math"
	"time"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for ; x < math.Pow(z, 2); {
		z += 0.1
	}
	return z
}

func day() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	default:
		fmt.Println("Too far away")
	}
}

func what_is_time() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	// fmt.Println(Sqrt(2))
	day()
	what_is_time()
	defer fmt.Println("world") // runs other code first, then this
	fmt.Println("hello")
}
