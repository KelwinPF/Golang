package main

import "fmt"
import "math"
type retangulo struct{
	altura float64
	largura float64 
}

type circulo struct{
	raio float64
}

type forma interface{
	area() float64
}

func escreverarea(f forma){
	fmt.Printf("a area e %0.2f ",f.area())
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio,2)
}

func main(){
	r:= retangulo{10,10}
	escreverarea(r)

	c:= circulo{10}
	escreverarea(c)
}