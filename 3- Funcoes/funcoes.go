package main
import "fmt"
func somar(n1 int8,n2 int8) int8{
	return n1+n2
}

func calculo(n1,n2 int8) (int8,int8){
	calculo1 := n1 + n2
	calculo2 := n1 - n2
	return calculo1,calculo2
}

func main(){
	soma:=somar(10,20)
	fmt.Println(soma)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto")
	fmt.Println(resultado)

	resultadoSoma,_ := calculo(19,20)
	fmt.Println(resultadoSoma)
}