package main

import (
	"fmt"
	"time"
)

func main() {


	y := 5
	switch y {
	case 5:
		fmt.Println("i = 5")
	case 10: 
		fmt.Println("i = 10")
	case 15: 
		fmt.Println("i = 15")
	}



	z := 10
	switch z {
	case 5:
		fmt.Println("z = 5")
	default:
		fmt.Println("z bilinmiyor")  // koşulların gerçekleşmediği durum
	}


	
	x := 5
	switch {
	case x == 5:
		fmt.Println("x = 5")
		fallthrough   // altındaki case i de çalıştır
	case x < 10:
		fmt.Println("x < 10")
	}



	switch x := 5; {
	case x == 5:
		fmt.Println("x 5 tir")
	case x < 10:
		fmt.Println("x 10 dan küçüktür")
	}



	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}



	switch time.Now().Weekday() { // bugün haftanın hangi günü
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default: 
		fmt.Println("It's a weekday")
	}

	fmt.Println(time.Now().Weekday())


	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}



	whatAmI := func(i any) {    // anonim fonksiyon. her türden veri kabul eder 
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am an int")
		default:
			fmt.Printf("dont know type %T\n", t)  // değişkenin veri tipi
 		}
	}
	

	whatAmI(true)
	whatAmI(5)
	whatAmI("hey")
}
