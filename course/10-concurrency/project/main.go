package main

import (
	"fmt"
	"price-calc-concurrency/prices"

	// "price-calc-concurrency/cmdmanager"
	"price-calc-concurrency/filemanager"
)


func main(){
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.5}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	
	// for every tax rate, go through prices and calculate price with tax
	for index, taxVal := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("results/result_%.0f.json", taxVal * 100))

		// cmdm := cmdmanager.New()
		// var cmdmPointer *cmdmanager.CMDManager = &cmdm

		priceJob := prices.NewTaxIncludedPriceJob(*fm, taxVal)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for i, _ := range doneChans {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}

	// for _, done := range doneChans {
	// 	<- done
	// }
}

// func printResults(results map[float64][]float64){
// 	for key, val := range results {
// 		fmt.Println(key, val)
// 	}
// }