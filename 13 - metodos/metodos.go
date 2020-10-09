package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}
func (u usuario) salvar(){
	fmt.Printf("usuario: %s salvo\n",u.nome)
}

func (u *usuario) niver(){
	u.idade++;
}

func (u usuario) demenor() bool {
	return u.idade<18
}

func main(){
	fmt.Println("metodos")

	user := usuario{"kelwin",25}
	user.salvar();

	demenor := user.demenor();
	fmt.Println(demenor)
	
	fmt.Println(user)
	user.niver();
	fmt.Println(user)
}