package main

import "fmt"

func closure() func(){
	texto := "dentro closure"

	funcao := func(){
		fmt.Println(texto)	
	}
	return funcao
}

func main(){
	texto :="na main"
	fmt.Println(texto)
	
	funcao := closure()
	funcao()
}