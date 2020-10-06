package main

import (
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
	"fmt"

)

func main(){
	fmt.Println("Ola mumdo")
	auxiliar.Escrever()

	erro:= checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
