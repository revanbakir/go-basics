package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func badServerHandler(w http.ResponseWriter, r *http.Request) {
	progress := 0
	for {
		if progress == 100 {
			fmt.Println("İndirme işlemi başarıyla tamamlandı.")
			break
		}
		time.Sleep(time.Second * 1)
		fmt.Println(fmt.Sprintf("İndirme işlemi devam ediyor. %%%d", progress))
		progress++
	}
	return
}

func main() {
	http.HandleFunc("/badserver", badServerHandler)
	fmt.Println("Listening 3010")
	log.Fatal(http.ListenAndServe(":3010", nil))
}

