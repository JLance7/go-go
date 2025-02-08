package main

import (
	"fmt"
	"os"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

var accountBalance float64

// func writeFloatToFile(fileName string, value float64){
// 	valueText := fmt.Sprint(value) // float to string
// 	os.WriteFile(fileName, []byte(valueText), 0644)
// }

// func getFloatFromFile(fileName string) (float64, error) {
// 	data, err := os.ReadFile(fileName)
// 	if err != nil{
// 		return 1000, errors.New("Failed to find file")
// 	}

// 	valueText := string(data) // convert to string
// 	value, err := strconv.ParseFloat(valueText, 64) // string to float
// 	if err != nil {
// 		return 1000, errors.New("Failed to parse stored value")
// 	}
// 	return value, nil
// }

func main(){
	start()
	fmt.Println("Welcome to the bank!")
	fmt.Println(randomdata.PhoneNumber())
	
	for {
		choice := options()
		runChoices(choice)
	}
}

func start(){
	var err error
	accountBalance, err = fileops.GetFloatFromFile("balance.txt")
	if err != nil {
		fmt.Println(err)
		// panic("Ahhhh")
	}
}

func options() int{
	cat_options()

	var choice int
	fmt.Scan(&choice)
	fmt.Println("Your choice:", choice)
	return choice
}

func looper(){
	for i := 1; i<=100; i++{
		fmt.Print(i, " ")
	}
}

func infiniteLooper(){
	i := 0
	for {
		fmt.Println(i)
		i += 1
	}
}

func booleanLooper(){
	var boolVal bool = true
	var inputVal int

	for boolVal == true {
		fmt.Println("true")
		fmt.Println("Continue? 1 for yes 2 for no")
		fmt.Scan(&inputVal)
		if inputVal == 2{
			boolVal = false
		}
	} 
	fmt.Println("false")
}

func runChoices(choice int){
	if choice == 1{
		fmt.Printf("Your account balance: %.2f\n", accountBalance)
	} else if choice == 2{
		var depositAmount float64

		fmt.Println("How much would you like to deposit?")
		fmt.Scan(&depositAmount)
		if (depositAmount < 0){
			fmt.Println("Deposit amout must be > 0")
		}
		accountBalance += depositAmount
		fileops.WiteFloatToFile("balance.txt", accountBalance)

		fmt.Println(depositAmount, "depositied")
	} else if choice == 3{
		var withdrawAmount float64

		fmt.Println("How much would you like to withdraw?")
		fmt.Scan(&withdrawAmount)
		if accountBalance - withdrawAmount >= 0 && withdrawAmount > 0{
			accountBalance -= withdrawAmount
			fileops.WiteFloatToFile("balance.txt", accountBalance)
			fmt.Println(withdrawAmount, "withdrawn")
		} else {
			fmt.Println("Can't withdraw that amount of money")
		}
	} else {
		fmt.Println("Exiting")
		os.Exit(0)
	}
}

func switching(input int){
	switch input{
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}
}

