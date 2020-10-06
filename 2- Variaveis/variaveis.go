package main

import "fmt"

func main(){
	var variavel1 string = "variavel"
	variavel2 := "variavel 2"

	var (
		variaveis3 string = "asdad"
		variaveis4 string = "dasd"
	)
	
	fmt.Println(variavel1,variavel2)
	variavel4,variavel6 := "testesd", "kkaaa"
	fmt.Println(variaveis3,variaveis4,variavel4,variavel6)
}