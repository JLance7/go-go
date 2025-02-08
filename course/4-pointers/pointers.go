package main

import (
	"fmt"
)

// use pointers to avoid unnecessary copying of values or to directly mutate values (directly manipulate value instead of needing to return a value)

func main(){
	i := 1
	fmt.Println(i)

	// iPointer := &i 
	var iPointer *int = &i;
	myFunc(iPointer)

	fmt.Println(i)
}

func myFunc(iPointer *int){
	// fmt.Println(*i)
	*iPointer = 2
}