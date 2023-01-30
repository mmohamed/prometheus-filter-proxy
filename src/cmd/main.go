package main

import (
	"os"

	proxy "github.com/mmohamed/prometheus-filter-proxy/src/proxy"
	"github.com/urfave/cli"
)

var (
	version = "dev"
)

func main() {
	app := cli.NewApp()
	app.Name = "Prometheus Filter & Authentication Proxy"
	app.Version = version
	app.Authors = []cli.Author{
		{Name: "Marouan MOHAMED", Email: "medmarouen@gmail.com"},
	}
	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Runs the Prometheus filter proxy",
			Action: proxy.Serve,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port",
					Usage: "Port to expose this prometheus proxy",
					Value: 9090,
				}, cli.StringFlag{
					Name:  "prometheus-server",
					Usage: "Prometheus server endpoint",
					Value: "http://localhost:9090",
				}, cli.StringFlag{
					Name:  "auth",
					Usage: "(Optional) username:password authentication info",
					Value: "",
					Required: false,
				},
			},
		},
	}
	app.Run(os.Args)
}
