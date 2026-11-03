package main

import "fmt"

func main() {

	d := [6]int{0,1,2,3,4,5}
	fmt.Println(d)
	var b []int = d[2:4]
	fmt.Println(b)


	e := [6]int{0,1,2,3,4,6}
	fmt.Println(e)
	var c []int = e[:3]
	fmt.Println(c)
	var f []int = e[4:]
	fmt.Println(f)


	x := [6]int{2,3,4,5,6,7}
	y := x[2:4]

	fmt.Println("x uzunluk: ", len(x))
	fmt.Println("x kapasite: ", cap(x))
	fmt.Println("x içerik: ", x)
	fmt.Println("y uzunluk: ", len(y))
	fmt.Println("y kapasite: ", cap(y))
	fmt.Println("y içerik: ", b) 

	var z []string 
	if z == nil {
		fmt.Println("boş")
	}
	
	z = append(z, "revan")
	z = append(z, "emir")
	fmt.Println(z)
	fmt.Println(len(z), cap(z))

	var isimler = []string{"revan", "emir", "melisa", "ali", "zeynep", "simay"}
	for m, n := range isimler {
		fmt.Printf("%d. index = %s\n", m, n)
	}
}
