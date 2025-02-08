// package main

// import "fmt"

// func main(){
// 	// 1) array that contains 3 hobbies
// 	arr := [3]string{"Archery", "Basketball", "Tennis"}
// 	fmt.Println(arr)

// 	// 2) output more data about array
// 	fmt.Println(arr[0])
// 	fmt.Println(arr[1])

// 	// 3) slice of first and second elements
// 	slice := arr[:2]
// 	fmt.Println(slice)
// 	slice = arr[0:2]
// 	fmt.Println(slice)

// 	// 4) second and last element
// 	slice = arr[1:]
// 	fmt.Println(slice)


// 	// 5) dynamic array
// 	dynamic := []string{"Get better", "Learn GO"}
// 	fmt.Println()
// 	for _, val := range dynamic {
// 		fmt.Println(val)
// 	}

// 	// 6) change second goal, add third goal
// 	dynamic[1] = "Learn Golang"
// 	dynamic = append(dynamic, "Make api")
// 	fmt.Println(dynamic)

// 	// 7) Create product struct and dynamic list of products 
// 	fmt.Println()
// 	type product struct{
// 		id int
// 		title string
// 		price float64
// 	}
// 	products := []product{
// 		{
// 			id: 1,
// 			title: "title1",
// 			price: 25.1,
// 		},
// 	}
// 	products = append(products, product{
// 		id: 2,
// 		title: "title2",
// 		price: 27.0,
// 	})
// 	fmt.Println(products)
// }