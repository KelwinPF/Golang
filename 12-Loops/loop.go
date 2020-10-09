package main

import "fmt"
import "time"

func main(){
	fmt.Println("Estrutura de controle")

	i := 0

	for i<10 {
		i++
		fmt.Println(i)
	}

	for j:=0;j<10;j++ {
		fmt.Println(j)
	}
	 
	nums := []int{5,2,5}
	for k,num := range nums{
		fmt.Println(k,num)
		time.Sleep(time.Second)
	}

	for l,letra := range "teste"{
		fmt.Println(l,letra,string(letra))
		time.Sleep(time.Second)
	}

	usuario := map[string] string {
		"nome": "pessoa",
		"sobrenome": "ribeiro",
	}
	
	for m,v := range usuario{
		fmt.Println(m,v)
	}
}