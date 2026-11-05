package main

import "fmt"


// generics aynı mantıkta çalışan bir fonksiyona farklı tipte veriler girebilmemizi sağlar

func arttır[n int | float64](sayı n) n {  // [...] içinde tipler jenerik tanımlanır
	return sayı + 1 
}



func main() {

	var sayı1 = 5
	var sayı2 = 4.6

	fmt.Println(arttır(sayı1))
	fmt.Println(arttır(sayı2))

}
