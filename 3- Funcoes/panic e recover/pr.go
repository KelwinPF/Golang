package main

import "fmt"

func recuperar(){
	if r:= recover(); r!=nil{
		fmt.Println("recuperei uhauhahu")
	}
}

func nota(n1,n2 float64)bool{
	defer recuperar();
	media:=(n1+n2)/2
	if media< 10{
		return true
	}else if media < 0{
		return false
	}

	panic("q media e essa")
}

func main(){
	fmt.Println(nota(12,11))
	fmt.Println("executa?")
}