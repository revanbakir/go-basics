package main

import "fmt"

func topla( a int, b int) int {
	return a + b
}

func topla2(k, l, m int) int {
	return k+l+m
}

func yazdir() {
	fmt.Println("yazı yaz")
}

func islem(sayi int) (x, y int) {
	x = sayi/2
	y = sayi*2
	return
}

func topla3(sayilar ...int) int { //variadic function: değişken sayıda argüman alan fonksiyon
	toplam := 0
	for _, n := range sayilar {
		toplam += n
	}
	return toplam
}


func faktoriyel(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * faktoriyel(a-1)
}

func main() {

	fmt.Println(topla(5, 3))
	yazdir()
	fmt.Println(islem(5)) // 2 10
	fmt.Println(topla2(1,2,3))
	fmt.Println(topla3(4,5,6,7,8,3))

	toplam := func(x, y int) int {
		return x + y
	}
	fmt.Println(toplam(2,3))

	fmt.Println(faktoriyel(3))
}


