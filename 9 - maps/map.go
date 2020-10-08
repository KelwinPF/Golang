package main

import "fmt"


func main(){
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":		"Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{"nome":		"Pedro",
		"sobrenome": "Silva",},
		"curso":{
			"nome":		"curso tal",
			"instituicao": "ufrn",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2,"nome")
	fmt.Println(usuario2)

	usuario2["emprego"] = map[string]string{
		"nome":"uber",
	}

	fmt.Println(usuario2)
}