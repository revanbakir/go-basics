package main

import "fmt"

func topla( a int, b int) int {

	return a + b
}

func yazdir() {
	fmt.Println("yazı yaz")
}

func islem(sayi int) (x, y int) {
	x = sayi/2
	y = sayi*2
	return
}


func main() {

	fmt.Println(topla(5, 3))
	yazdir()
	fmt.Println(islem(5)) // 2 10
}


