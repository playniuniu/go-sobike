package main

import (
	"bikelib"
	"fmt"
	"maplib"
	"os"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
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

type chanStruct struct {
	Error error
	Data  interface{}
	Type  string
}

func initCLI() {
	app := cli.NewApp()
	app.Name = "SoBike"
	app.Version = appVersion
	app.Usage = "Search bike around you"
	cli.AppHelpTemplate = fmt.Sprintf("%s%s", appBanner, cli.AppHelpTemplate)

	var argsCity string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "city, c",
			Usage:       "Set address `CITY`",
			Destination: &argsCity,
		},
	}

	app.Action = func(c *cli.Context) error {
		var argsAddress string
		if c.NArg() > 0 {
			argsAddress = c.Args()[0]
			run(argsAddress, argsCity)
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	app.Run(os.Args)
}

func run(address string, city string) {
	color.HiCyan(appBanner)

	mapChan := make(chan chanStruct, 1)
	go getMapData(mapChan, address, city)
	mapRes := <-mapChan
	close(mapChan)

	if mapRes.Error != nil {
		return
	}

	bikeChan := make(chan chanStruct, 2)
	mapLoc := mapRes.Data.(maplib.MapLocation)
	color.HiCyan("正在为你寻找 %v 附近的自行车", mapLoc.Address)

	go getMobikeData(bikeChan, mapLoc)
	go getOfoData(bikeChan, mapLoc)

	for i := 0; i < cap(bikeChan); i++ {
		select {
		case dataChan := <-bikeChan:
			displayBikeData(dataChan, mapLoc)
		}
	}
	close(bikeChan)
}

func getMobikeData(bikeChan chan chanStruct, mapLoc maplib.MapLocation) {
	bike := bikelib.Mobike{
		Lat:      mapLoc.Lat,
		Lng:      mapLoc.Lng,
		CityCode: mapLoc.CityCode,
	}
	data, err := bike.GetNearbyCar()
	if err != nil {
		bikeChan <- chanStruct{
			Error: err,
			Data:  nil,
			Type:  "mobike",
		}
	} else {
		bikeChan <- chanStruct{
			Error: nil,
			Data:  data,
			Type:  "mobike",
		}
	}
}

func getOfoData(bikeChan chan chanStruct, mapLoc maplib.MapLocation) {
	bike := bikelib.Ofobike{
		Lat: mapLoc.Lat,
		Lng: mapLoc.Lng,
	}
	data, err := bike.GetNearbyCar()
	if err != nil {
		bikeChan <- chanStruct{
			Error: err,
			Data:  nil,
			Type:  "ofo",
		}
	} else {
		bikeChan <- chanStruct{
			Error: nil,
			Data:  data,
			Type:  "ofo",
		}
	}
}

func getMapData(mapChan chan chanStruct, address string, city string) {
	mapObj := maplib.MapAddr{
		Address: address,
		City:    city,
	}

	mapLoc, err := mapObj.GetGeoLoc()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get map data")
		mapChan <- chanStruct{
			Error: err,
			Data:  nil,
			Type:  "map",
		}
	} else {
		mapChan <- chanStruct{
			Error: nil,
			Data:  mapLoc,
			Type:  "map",
		}
	}
}

func displayBikeData(bikeChan chanStruct, mapLoc maplib.MapLocation) {
	if bikeChan.Type == "mobike" {
		bikeList := bikeChan.Data.([]bikelib.MobikeCar)
		screen := color.New(color.FgHiMagenta)
		screen.Printf("\n---------------------\n")
		screen.Printf("摩拜自行车, 共 %v 辆\n", len(bikeList))
		screen.Printf("---------------------\n")
		for _, el := range bikeList {
			screen.Printf("车牌号: %v, 距离您: %v 米\n", el.DistID, el.Distance)
		}
	}

	if bikeChan.Type == "ofo" {
		bikeList := bikeChan.Data.([]bikelib.OfoCar)
		screen := color.New(color.FgHiYellow)
		screen.Printf("\n---------------------\n")
		screen.Printf("Ofo 自行车, 共 %v 辆\n", len(bikeList))
		screen.Printf("---------------------\n")
		for _, el := range bikeList {
			distance := int(maplib.GeoDistance(mapLoc.Lat, mapLoc.Lng, el.Lat, el.Lng))
			screen.Printf("车牌号: %v, 距离您: %v 米\n", el.Carno, distance)
		}
	}
}

func main() {
	initCLI()
}
