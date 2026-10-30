package main

import "fmt"

func main() {

	defer fmt.Println("İlk cümle")    // LIFO(son giren ilk çıkar) -> stack mantığı
	defer fmt.Println("ikinci cümle")
	defer fmt.Println("üçüncü cümle")
	defer fmt.Println("dördüncü cümle")
	fmt.Println("beşinci cümle")



	fmt.Println("sayıyor...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("bitti")
}

