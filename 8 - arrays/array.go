package main

import "fmt"
import "reflect"


func main(){

	fmt.Println("array")

	var array1[5] int
	array1[0] = 2
	fmt.Println(array1)

	array2 := [5]int{1,2,3,4,5} 
	fmt.Println(array2)

	array3 := [...]int{1,2,4,5,2,1}
	fmt.Println(array3)
	
	//slice
	slice := []int{21,2}
	fmt.Println(slice)
	slice = append(slice,18)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 55
	fmt.Println(slice2)

	//array interno

	slice3 := make([]float32,10,11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))  //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3,18)
	fmt.Println(slice3)
	fmt.Println(len(slice3))  //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3,22)
	fmt.Println(slice3)
	fmt.Println(len(slice3))  //tamanho
	fmt.Println(cap(slice3)) //capacidade
}