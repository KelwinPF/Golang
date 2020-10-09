package main

import "fmt"

func main(){

	retorn := func(texto string) string{
		return fmt.Sprintf("recebido -> %s",texto)
	}("parameter")
	fmt.Println(retorn)

}