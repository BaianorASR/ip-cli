package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI IPs search"
	app.Usage = "Search IP addresses and Name from web server"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP addresses from web server",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "server",
			Usage:  "Search name of web server",
			Flags:  flags,
			Action: searchServes,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServes(c *cli.Context) {
	host := c.String("host")

	names, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, name := range names {
		fmt.Println(name.Host)
	}

}
