package main

import "fmt"

func main() {


	a := 8
	fmt.Println(&a) 

	b := &a // a nın adresi - b
	fmt.Println(b)  
	fmt.Println(*b) // b adresindeki değer - a 

	c := 6
	ekle(&c)
	fmt.Println(c)
}


func ekle(v *int) {
	 *v += 5
}
