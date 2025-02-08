package main

import (
	"fmt"
	"math"
)

func Calculate(){
	var investmentAmount float64 // default null value given which is 0.0
	var years float64 = 10
	expectedReturnRate := 5.5
	const inflationRate = 2.5

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Amoount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := float64(investmentAmount) * math.Pow(1 + expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years);
	//fmt.Println(futureValue)
	fmt.Printf("Future Value %.2f\n", futureValue)

	// fmt.Println(futureRealValue)
	fmt.Printf("Future real value %.2f\n", futureRealValue)

	formattedFV := fmt.Sprintf(`Future value: %.1f
	awesome
	`, futureValue)
	fmt.Print(formattedFV)
}

/*
int
float64
string in double quotes
bool

uint
int32
rune (int32 ascii unicode character)
uint32
int64
*/