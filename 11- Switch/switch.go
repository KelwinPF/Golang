package main

import "fmt"

func diasemana(numero int) string{
	switch numero{
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"	
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "incorreto"
	}
}

func diasemana2(numero int) string{
	var dia string
	switch {
	case numero == 1:
		dia = "domingo"
		//fallthrough
	case numero == 2:
		dia = "segunda"
	case numero == 3:
		dia = "terça"
	case numero == 4:
		dia = "quarta"	
	case numero == 5:
		dia = "quinta"
	case numero == 6:
		dia = "sexta"
	case numero == 7:
		dia = "sabado"
	default:
		dia = "incorreto"
	}

	return dia
}

func main(){
	fmt.Println("switch")

	dia := diasemana(3)
	fmt.Println(dia)	
	
	dia2 := diasemana2(4)
	fmt.Println(dia2)	
}