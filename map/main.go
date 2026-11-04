package main

import(
"fmt"
"maps"
)


type insan struct {
	kisi1, kisi2, kisi3 string
}

var m = map[string]insan {
	"erkekler": insan{"ali", "efe", "berk"},
	"kadınlar": insan{"revan", "nil", "idil"},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["erkekler"])
	fmt.Println(m["kadınlar"])




	m1 := make(map[string]int)   // make(map[key-type]val-type)
	m1["k1"] = 7
	m1["k2"] = 13

	fmt.Println("map: ", m1)
	v1 := m1["k1"]
	fmt.Println("v1: ", v1)
	v3 := m1["k3"]
	fmt.Println("v3: ", v3)  // 0
	fmt.Println("len: ", len(m1))


	delete(m1, "k2")
	fmt.Println("map: ", m1)

	clear(m1)
	fmt.Println("map: ", m1)

	_, prs := m1["k2"]
	fmt.Println("prs: ", prs)


	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
	
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
 	}

}
