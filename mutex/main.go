package main

import (
"fmt"
"sync"  // mutexi kullanmak için
)


// global mutex nesnesi oluşturma
var mt sync.Mutex


func paraCek(bakiye *float64, cekilecekMiktar float64, wg *sync.WaitGroup) {


	/* mt isimli mutexi bu işlem yapılırken kitliyoruz
	bu sayede mt mutexini başka işlemler kullanamıyor */

	mt.Lock()

	// asenkron olmasını istemediğimiz işlem
	*bakiye -= cekilecekMiktar

	fmt.Printf("yeni bakiye: %2.f\n", *bakiye)


	/* diğer işlemlerin de kullanılabilmesi için mutexi açalım
	mt mutex açılınca diğer asenkron işlemdeki mt mutexi çalışmaya başlar */

	mt.Unlock()
	fmt.Println("para çekme işlemi tamamlandı")

	/* waitgroup ile işlem tamamlandı
	wg sayacı 2 den 1 e düştü */

	wg.Done()
}

func paraYatir(bakiye *float64, yatirilanMiktar float64, wg *sync.WaitGroup) {

	mt.Lock()
	*bakiye += yatirilanMiktar
	fmt.Printf("yeni bakiye: %2.f\n", *bakiye)
	mt.Unlock()
	fmt.Println("yatirma işlemi tamamlandı")
	wg.Done()
}





func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	var bakiye float64 = 100
	fmt.Printf("ilk bakiye: %.2f\n", bakiye)
	
	go paraCek(&bakiye, 25, &wg)
	go paraYatir(&bakiye, 65, &wg)

	wg.Wait()
}
