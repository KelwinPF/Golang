package main

import "fmt"


func main(){

	var vari1 int = 10
	var vari2 int = vari1

	fmt.Println(vari1,vari2)

	vari2++;

	fmt.Println(vari1,vari2)

	//ponteiro
	var vari3 int = 100
	var ponteiro *int = &vari3

	fmt.Println(vari3,*ponteiro)
}