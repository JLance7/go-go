package lists

import (
	"fmt"
)

// func main(){
// 	prices := [4]float64{1, 2, 3, 4}
// 	fmt.Println(prices)

// 	var products [4]float64 
// 	products = [4]float64{2.1, 2.2, 2.4, 2.8}
// 	fmt.Println(products[0])

// 	// slice
// 	featuredPrices := products[1:3]
// 	fmt.Println(featuredPrices)

// 	morePrices := products[:3]
// 	fmt.Println(morePrices)
// 	fmt.Println(len(morePrices))

// 	highLighetedPrices := products[:1]
// 	fmt.Println(len(highLighetedPrices), cap(highLighetedPrices))
// }

func main(){
	// dynamic slice like array with dynamic length
	prices := []float64{} // when resizing makes a new array behind the scenes
	prices = append(prices, 1.1)
	prices = append(prices, 1.2, 1.3)
	fmt.Println(prices)
	fmt.Println()
	for i := 0; i<len(prices); i++ {
		fmt.Println(prices[i])
	}

	for index, value := range prices {
		fmt.Println("Index:", index, "Value:", value)
	}

	// merging slices
	prices = []float64{4.5, 4.6}
	discountPrices := []float64{5.5, 5.6}
	prices = append(prices, discountPrices...) // unpacking list values
	fmt.Println(prices)
}

