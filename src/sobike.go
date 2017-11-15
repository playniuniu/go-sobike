package main

import (
	"bikelib"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const appVersion = "v0.1"
const appBanner = `
███████╗ ██████╗ ██████╗ ██╗██╗  ██╗███████╗
██╔════╝██╔═══██╗██╔══██╗██║██║ ██╔╝██╔════╝
███████╗██║   ██║██████╔╝██║█████╔╝ █████╗  
╚════██║██║   ██║██╔══██╗██║██╔═██╗ ██╔══╝  
███████║╚██████╔╝██████╔╝██║██║  ██╗███████╗
╚══════╝ ╚═════╝ ╚═════╝ ╚═╝╚═╝  ╚═╝╚══════╝

`

func initCLI() *cli.App {
	color.Cyan(appBanner)
	app := cli.NewApp()
	app.Name = "SoBike"
	app.Version = appVersion
	app.Usage = "Search bike around you"

	return app
}

func printBikeData() {
	bike := bikelib.Ofobike{Lat: 123.0, Lng: 345.0}
	data, err := bike.GetNearbyCar()
	if err != nil {
		return
	}
	for _, el := range data {
		fmt.Println(el)
	}
}

func main() {
	app := initCLI()
	app.Action = func(c *cli.Context) error {
		return nil
	}
	app.Run(os.Args)
	printBikeData()
}
