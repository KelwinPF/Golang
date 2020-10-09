package main
import "fmt"

func fibo(pos uint) uint{
	if pos <= 1 {
		return pos
	}

	return fibo(pos-2) + fibo(pos-1)
}


func main(){
	fmt.Println("recursivas")
	posicao := uint(10)
	for i := uint(0); i<posicao;i++{
		fmt.Println(fibo(i+1))
	}
}