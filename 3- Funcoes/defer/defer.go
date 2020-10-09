package main

import "fmt"

func funcao1(){
	fmt.Println("f1")	
}

func funcao2(){
	fmt.Println("f2")
}

func main(){
	defer funcao1()

	funcao2()
}