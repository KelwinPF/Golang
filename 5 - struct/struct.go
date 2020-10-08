package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}
func main(){

	fmt.Println("struct")
	var u usuario
	u.nome = "Daaav"
	u.idade = 21
	fmt.Println(u)

	adress := endereco{"rua x",99}

	usuario2 := usuario{"user",21,adress}
	fmt.Println(usuario2)
	
	usuario3 := usuario{nome:"user"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade:22}
	fmt.Println(usuario4)
}