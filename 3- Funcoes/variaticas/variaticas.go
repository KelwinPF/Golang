package main
import "fmt"

func soma(ns ...int) int {
	soma := 0 
	for _, n := range ns{
		soma += n
	}
	return soma
}


func main(){
	somar := soma(10,20,30,40,50)
	fmt.Println(somar) 

}