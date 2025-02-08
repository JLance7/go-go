package main

import (
	"fmt"
	"price-calculator/cmdmanager"
	// "price-calculator/filemanager"
	"price-calculator/prices"
)


func main(){
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.5}
	
	// for every tax rate, go through prices and calculate price with tax
	for _, taxVal := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxVal * 100))
		cmdm := cmdmanager.New()
		var cmdmPointer *cmdmanager.CMDManager = &cmdm
		priceJob := prices.NewTaxIncludedPriceJob(*cmdmPointer, taxVal)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}

// func printResults(results map[float64][]float64){
// 	for key, val := range results {
// 		fmt.Println(key, val)
// 	}
// }