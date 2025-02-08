package main

import (
	"fmt"
	"errors"
	"os"
)

func main(){
	// input revenue, taxes, & tax rate
	// calculate earning before and after tax and ratio and outputs 3 values

	var revenue float64
	var expenses float64
	var tax_rate float64
	var err error

	revenue, err = getUserInput("Revenue: ")
	if err != nil{
		panic("revenue input error")
	}

	expenses, err = getUserInput("Expenses: ")
	if err != nil{
		panic("expenses input error")
	}

	tax_rate, err = getUserInput("Tax Rate: ")
	if err != nil{
		panic("tax rate input error")
	}

	earning_before_tax := ebt(revenue, expenses)
	after_tax := after_tax(earning_before_tax, tax_rate)
	ratio := ratio(earning_before_tax, after_tax)

	storeCalculatedResults(earning_before_tax, after_tax, ratio)
	fmt.Println("\nEarning before tax: ", earning_before_tax)
	fmt.Println("After tax: ", after_tax)
	fmt.Println("Ratio: ", ratio)
}

func getUserInput(input string) (float64, error){
	var userInput float64
	fmt.Println(input)
	fmt.Scan(&userInput)
	if userInput <= 0{
		return 0, errors.New("Input must be greater than 0")
	}
	return userInput, nil
}

func ebt(revenue float64, expenses float64) float64 {
	ebt := revenue - expenses
	return ebt
}

func after_tax(earning_before_tax float64, tax_rate float64) float64{
	after_tax := earning_before_tax - (earning_before_tax * (tax_rate/100))
	return after_tax
}

func ratio(earning_before_tax float64, after_tax float64) float64{
	return earning_before_tax / after_tax
}

func storeCalculatedResults(earning_before_tax float64, after_tax float64, ratio float64){
	var results string
	// float to string
	ebtText := fmt.Sprint(earning_before_tax)
	afterTaxText := fmt.Sprint(after_tax)
	ratioText := fmt.Sprint(ratio)

	results += "EBT: " + ebtText + "\n"
	results += "After tax: " + afterTaxText + "\n"
	results += "Ratio: " + ratioText
	storeToFileResults(results, "results.txt")
}

func storeToFileResults (results string, file string){
	// float to string
	os.WriteFile(file, []byte(results), 0644)
}