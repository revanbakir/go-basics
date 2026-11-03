package main

import "fmt"

func main() {

	var a [3]string
	a[0] = "revan"
	a[1] = "ali efe"
	a[2] = "nil"

	fmt.Println(a)
	fmt.Println(a[1])


	b := [3]string{"revan", "ali efe", "nil"}
	fmt.Println(b)


	var c [5]int 
	fmt.Println("emp:", c)  // hepsi 0
	c[4] = 100
	fmt.Println("set:", c)
	fmt.Println("get:", c[4])
	fmt.Println("len:", len(c))



	d := [5]int{1,2,3,4,5}
	fmt.Println(d)

	d = [...]int{1,2,3,4,5}
	fmt.Println(d)

	d = [...]int{100, 3: 400, 500}  // {10, 0, 0, 400, 500}
	fmt.Println(d)


	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)



	twoD = [2][3] int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("2d: ", twoD)

}
