package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App{
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}

	app := cli.NewApp()
	app.Name = "App Go Comand line"
	app.Usage = "Search IPs and server names on Internet"
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Buscando Ips de endereço na internet",
			Flags: flags,
			Action: BuscarIps,
		},
		{
			Name: "server",
			Usage: "Buscando servidores hospedados na internet",
			Flags: flags,
			Action: BuscarServidor,
		},

	}
	return app
}

func  BuscarIps(c *cli.Context)  {
	host := c.String("host")
	
	ips, erros := net.LookupIP(host)
	if erros != nil {
		log.Fatal(erros)
	}
	
	for _, i := range ips {
		fmt.Println(i)
	}
}

func BuscarServidor(c *cli.Context) {
	host := c.String("host")

	servers, erros := net.LookupNS(host) // name server
	if erros != nil {
		log.Fatal(erros)
	}

	for _, ss := range servers {
		fmt.Println(ss)
	}
}