package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	var entrada string
	//CAPTURA STRINGS
	fmt.Print("Ingrese la cantidad de strings a capturar:")
	fmt.Scan(&n)
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		fmt.Print("String: ")
		fmt.Scan(&entrada)
		s = append(s, entrada)
	}
	//ORDENA DE MANERA ASCENDENTE
	fmt.Println(s)  //slice antes de ordenar
	sort.Strings(s) //ordena de manera ascendente
	fmt.Println(s)  //slice despues de ordenar
	//crea archivo
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//guarda en archivo
	for _, s2 := range s {
		file.WriteString(s2 + "\n")
	}
	//ORDENA DE MANERA DESCENDENTE
	sort.Sort(sort.Reverse(sort.StringSlice(s))) //ordena de mañera descendente
	fmt.Println(s)                               //slice despues de ordenar
	//crea archivo
	file2, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()
	//guarda en archivo
	for _, s2 := range s {
		file2.WriteString(s2 + "\n")
	}
}
