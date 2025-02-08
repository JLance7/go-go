package main

import (
	"fmt"
)

var global_var int = 5

func main(){
	callFunc("hello")
	fmt.Println(returnResult(1))
	fmt.Println(global_var)
	val1, val2 := returnValues()
	fmt.Println(val1, val2)
}

func callFunc(arg string){
	fmt.Println("called", arg)
}

func returnResult(num int) int{
	return num + 1
}

func returnValues() (int, int) {
	val := 1
	val2 := 2
	return val, val2
}

func returnValues2() (val int, val2 int) {
	val = 1
	val2 = 2
	return val, val2
}