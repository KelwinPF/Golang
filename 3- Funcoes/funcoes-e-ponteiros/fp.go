package main

import "fmt"

func invsinal(num int) int {
	return num * -1
}

func invsinalponteiro(num *int){
	*num = *num * -1
}
func main(){
	numero := 10
	inv := invsinal(numero)
	fmt.Println(numero)
	fmt.Println(inv)

	numero2 := 20
	fmt.Println(numero2)
	invsinalponteiro(&numero2)
	fmt.Println(numero2)
}