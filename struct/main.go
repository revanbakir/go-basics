package main

import "fmt"


// go da clas lar yoktur onun yerine struct(yapı) kullanılır

type kisi struct {  // type ile yeni bir tür. bu tür de struct   

	isim string
	soyisim string
	yas int
}


type insan struct {

	isim string
	yas int
}


func (i insan) tanıt() {
	fmt.Printf("Merhaba, Ben %s, %d yaşındayım.", i.isim, i.yas) //%s-string türlü veriler, %d-integer türlü veriler
}



func main () {

	kisi1 := kisi{"kaan", "kuşçu", 25}
	fmt.Println(kisi1)

	kisi1.isim, kisi1.soyisim = "revan", "bakır"
	fmt.Println(kisi1)


	kisi2 := kisi{}
	kisi2.isim, kisi2.soyisim = "melisa", "duran"
	kisi2.yas = 30
	fmt.Println(kisi2)

	kisi3 := kisi{soyisim: "duman", isim: "selin", yas: 24}
	fmt.Println(kisi3)


	person := struct {    // structı bir değişken içine tanımladık. 
		ad, soyad string
		}{"Kemal", "Ataturk"}
		fmt.Println(person)

	kisi4 := insan{"liva", 23}
	kisi4.tanıt()	
}

	
