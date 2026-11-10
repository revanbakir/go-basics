package main

import "fmt"

// bazen programın açılışında mainden önce ayarlamamız gereken önkoşullar içi init() fonksiyonu kullanılır

var degisken = "revan"  // global değişken

func init() {
	degisken = "melisa"
	fmt.Println("degisken init: ", degisken)
}


func main() {

	degisken = "ali efe"
	fmt.Println("deisken main: ", degisken)  // global tanımlanan değişkeni init() içind edeğiştirdiğimizde farklı fonskiyonlarda da kullanabiliriz
}

