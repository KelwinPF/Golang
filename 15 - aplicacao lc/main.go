package main

import "fmt"
import "linha-de-comando/app"
import "os"
import "log"

func main(){
	fmt.Println("inicio")

	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}