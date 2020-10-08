package main

import "fmt"


func main(){
	fmt.Println("Estrutura de controle")

	numero := 20

	if numero > 25{
		fmt.Println("maior")
	}else {
		fmt.Println("menor ou igula")
	}

	if num2 := numero; num2 > 0 {
		fmt.Println("maior q 0")
	}else if num2 == 0 {
		fmt.Println("maior q 0")
	}else{
		fmt.Println("menor q 0")
	}

}