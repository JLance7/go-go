package main

import "fmt"

// type alias
type floatMap map[string]string

func (m floatMap) output(){
	fmt.Println(m)
}

func main(){
	// make go aware that we will add 2 elements for efficiency
	userNames := make([]string, 2, 5) // slice with 2 null (empty string values), third arg of 5 for capcity can be supplied to allocate that much memory space initially (rough estimate given)
	fmt.Println(userNames)

	// userNames = append(userNames, "Max")
	// userNames = append(userNames, "John")

	userNames[0] = "Max"
	userNames[1] = "John"

	fmt.Println(userNames)

	courses := make(floatMap, 3)
	courses["title"] = "course1"
	courses["title2"] = "course2"
	courses.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, val := range courses {
		fmt.Println(key, val)
	}
}