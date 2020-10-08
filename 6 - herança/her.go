package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main(){

	fmt.Println("herança")
	
	p1 := pessoa{"kelwin","ferreira",25,178}
	fmt.Println(p1)

	e1 := estudante{p1,"engenharia","UFRN"}
	fmt.Println(e1.altura)
}