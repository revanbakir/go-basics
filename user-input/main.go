package main

import "fmt"

func main() {

	var first string
        var last string

	fmt.Println("enter your first name: ")
	fmt.Scanln(&first)

	fmt.Println("enter your last name: ")
	fmt.Scanln(&last)

	fmt.Println("your full name is: ", first, " ", last)


	var kelime1, kelime2 string
	fmt.Println("2 kelime gir: ")

	fmt.Scanf("%s %s", &kelime1, &kelime2)
	fmt.Println(kelime1)
	fmt.Println(kelime2)

}
