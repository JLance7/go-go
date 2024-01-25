package main

import "fmt"

var (
	v2 = Vertex{5, 10}
	global_p = &v2
)

func main(){
	i, j := 42, 2701
	fmt.Printf("i: %v, j: %v\n", i, j)
	var p *int = nil
	p = &i
	fmt.Println(*p)

	p2 := &j
	fmt.Println(*p2)
	fmt.Println(Vertex{1, 2})

	v := Vertex{2, 4}
	fmt.Printf("x: %v, y: %v", v.X, v.Y)
	p3 := &v
	p3.X = 2e9
	fmt.Println(v)

	fmt.Println(*global_p)

	var a [2]string //array
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	primes := [6]int{2, 3, 5, 7, 11, 13}
	for i := 0; i < len(primes); i++ {
		fmt.Print(i)
	}
	fmt.Println()
	//slice
	var s []int = primes[1:4] // 1 to 3
	fmt.Println(s)

	more_arrays()
	more_slices()
}

func more_arrays(){
	names := []string{
		"Josh",
		"Mike",
		"alex",
	}

	fmt.Println(names)
	a := names[0:2]
	fmt.Printf("a: %v, T: %T", a, a)

	r := []bool{true, false, true}
	fmt.Println(r)
	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}
	fmt.Println(s)
	a2 := [10]int{1, 2, 3, 4, 5}
	fmt.Println(a2[:3])
}

func more_slices(){
	s := []int{2, 3, 4, 5, 6, 7}
	printSlice(s)
	s2 := s[:1]
	printSlice(s2)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type Vertex struct {
	X int
	Y int
}