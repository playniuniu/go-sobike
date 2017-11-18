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

type chanData struct {
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
	// 打印 Banner
	color.HiCyan(appBanner)

	// 获取地图的 Lng, Lat
	mapChan := make(chan chanData, 1)
	go getMapData(mapChan, address, city)
	mapRes := <-mapChan
	close(mapChan)

	if mapRes.Error != nil {
		return
	}
	mapLoc := mapRes.Data.(maplib.MapLocation)

	// 获取单车信息
	var bike bikelib.BikeInterface
	bikeChan := make(chan chanData, 2)
	color.HiCyan("正在为你寻找 %v 附近的自行车", mapLoc.Address)

	// 获取 Mobike
	go func() {
		bike = bikelib.Mobike{
			Lat:      mapLoc.Lat,
			Lng:      mapLoc.Lng,
			CityCode: mapLoc.CityCode,
		}
		getBikeData(bikeChan, bike, "mobike")
	}()

	// 获取 Ofo
	go func() {
		bike = bikelib.Ofobike{
			Lat: mapLoc.Lat,
			Lng: mapLoc.Lng,
		}
		getBikeData(bikeChan, bike, "ofo")
	}()

	// 打印单车信息
	for i := 0; i < cap(bikeChan); i++ {
		select {
		case dataChan := <-bikeChan:
			displayBikeData(dataChan, mapLoc)
		}
	}
	close(bikeChan)
}

func getMapData(mapChan chan chanData, address string, city string) {
	mapObj := maplib.MapAddr{
		Address: address,
		City:    city,
	}

	mapLoc, err := mapObj.GetGeoLoc()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get map data")
		mapChan <- chanData{
			Error: err,
			Data:  nil,
			Type:  "map",
		}
	} else {
		mapChan <- chanData{
			Error: nil,
			Data:  mapLoc,
			Type:  "map",
		}
	}
}

func getBikeData(bikeChan chan chanData, bike bikelib.BikeInterface, bikeType string) {
	data, err := bike.GetNearbyCar()
	if err != nil {
		bikeChan <- chanData{
			Error: err,
			Data:  nil,
			Type:  bikeType,
		}
	} else {
		bikeChan <- chanData{
			Error: nil,
			Data:  data,
			Type:  bikeType,
		}
	}
}

func displayBikeData(bikeChan chanData, mapLoc maplib.MapLocation) {
	bikeList := bikeChan.Data.([]bikelib.BikeData)
	var screen *color.Color

	if bikeChan.Type == "mobike" {
		screen = color.New(color.FgHiMagenta)
		screen.Printf("\n---------------------\n")
		screen.Printf("摩拜自行车, 共 %v 辆\n", len(bikeList))
		screen.Printf("---------------------\n")
	} else {
		screen = color.New(color.FgHiYellow)
		screen.Printf("\n---------------------\n")
		screen.Printf("Ofo 自行车, 共 %v 辆\n", len(bikeList))
		screen.Printf("---------------------\n")
	}

	for _, el := range bikeList {
		distance := maplib.GeoDistance(mapLoc.Lng, mapLoc.Lat, el.Lng, el.Lat)
		screen.Printf("车牌号: %v, 距离您: %v 米\n", el.CarNo, int(distance))
	}
}

func main() {
	initCLI()
}
