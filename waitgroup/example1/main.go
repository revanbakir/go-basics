package main 

import (
"fmt"
"time"
"sync"
)

func main() {

	var wg sync.WaitGroup // waitgroupu kullanmak için bir waitgroup nesnesi. sayaç kutusu
	wg.Add(2) // 2 işlem yapılacağını belirttik


	go fonksiyon1(&wg) // fonksiyonlara oluşturduğumuz wg örneğinin bellekteki adresini veeriyoruz
	go fonksiyon2(&wg)

	fmt.Println("merhaba dünya")

	wg.Wait()  // asenkron işlemlerin bitmesini bekliyoruz yani waitgroup un 0 a düşmesini. program beklemeden kendini sonlandırmayacak

	fmt.Println("waitgrouplar tamamlandı")

}


func fonksiyon1(wg *sync.WaitGroup) { // waitgroup nesnesini pointer olarak veriyoruz 

	time.Sleep(2*time.Second)
	fmt.Println("fonk1 tamamlandı")

	wg.Done() // waitgroup nesnesini 1 azalttık
}


func fonksiyon2(wg *sync.WaitGroup) {
	time.Sleep(2*time.Second)
	fmt.Println("fonk2 tamamlandı")

	wg.Done()
}
	



