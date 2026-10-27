package main

import (
 "fmt"
 "math"
)

const s string = "constant"
// const ile := beraber kullanılmaz
// const const s = "constant" diye de yazabiliriz

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)


	fmt.Println(s)
	const n = 500000000
	const d = 3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	fmt.Println(&f) //değişkenin bellekteki adresi

	i := 42
	m := float64(i)
	u := uint(m)
	fmt.Println(i,m,u)


	deneme := string(8378)
	fmt.Println(deneme) // sayıyı strinnge dönüştürdüğümüzde karakter olarak değer alır
}

