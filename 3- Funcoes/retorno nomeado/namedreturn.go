package main
import "fmt"

func somasub(n1 int,n2 int) (soma int,sub int){
	soma = n1+n2
	sub = n1-n2
	return
}


func main(){
	soma,sub:=somasub(10,20)
	fmt.Println(soma,sub)

}