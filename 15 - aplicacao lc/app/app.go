package app

import "github.com/urfave/cli"
import "fmt"
import "log"
import "net"
//vai retornar a aplicacao de linha de comando para ser executada
func Gerar() *cli.App {
	app:= cli.NewApp()
	app.Name = "aplicacao linha de comando"
	app.Usage = "busca de ips e nomes de servers"
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "busca de ips na internet",
			Flags: flags,
			Action: buscarips,
		},
		{
			Name: "servidores",
			Usage: "busca nome de servidores na internet",
			Flags: flags,
			Action: buscarservers,
		},
	}
	
	return app
}

func buscarservers(c *cli.Context){
	host := c.String("host")

	servers,erro := net.LookupNS(host)
	if erro !=nil{
		log.Fatal(erro)
	}
	for _,server := range servers{
		fmt.Println(server)
	}
}

func buscarips(c *cli.Context){
	host := c.String("host")

	ips,erro := net.LookupIP(host)
	if erro !=nil{
		log.Fatal(erro)
	}

	for _,ip := range ips{
		fmt.Println(ip)
	}
}