yazdir() (no value) used as valueyazdir() (no value) used as valuepackage main

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

func sum (nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum: ", total)
}


func faktoriyel(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * faktoriyel(a-1)
}


func vals() (int, int) {
	return 3, 7
}


func yazdir2() func(s string) {  // fir fonksiyon, fonksiyon tipinde veri döndürüyor
	return func(s string) {
		fmt.Println(s)
	}
}

func fonksiyonumuz (girdi int) (int, int) {
	islem1 := girdi/2
	islem2 := girdi/4
	return islem1, islem2
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


	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()  // boş tanımlayıcı
	fmt.Println(c)


	sum(1,3,5)
	sum(1,3)
	nums := []int{1,2,3,4}
	sum(nums...)

	toplam3 := func(a, b int) int {     // Closure(atanmış) fonksiyon: değişkenleri fonksiyon olarak tanımlayabiliriz
		return a + b
	}
	fmt.Println(toplam3(4,5))


	p := yazdir2()
	p("Hello world")

	metin := "merhaba dünya"
	func (a string) {
		fmt.Println(a)
	}(metin)    // anonim fonksiyon


	_, dördeböl := fonksiyonumuz(16)  // kullanmak istemediğimiz ögeri boş tanımladık
	fmt.Println(dördeböl)

}


