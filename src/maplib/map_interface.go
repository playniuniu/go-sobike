package maplib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const apiURL = "http://restapi.amap.com/v3/geocode/geo"
const apiKey = "ef5655ca17a2c9d6adf67810b12cf9c1"

// MapAddr struct
type MapAddr struct {
	Address string
	City    string
}

// GetGeoLoc function
func (mapaddr *MapAddr) GetGeoLoc() (MapLocation, error) {
	var res MapLocation

	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("address", mapaddr.Address)
	if mapaddr.City != "" {
		params.Add("city", mapaddr.City)
	}

	geoData, err := HTTPGet(apiURL, params)
	if err != nil {
		return res, err
	}

	return mapaddr.parseJSON(geoData)
}

func (mapaddr *MapAddr) parseJSON(jsonData []byte) (MapLocation, error) {
	var res MapLocation
	var parseData GaodeJSON

	err := json.Unmarshal(jsonData, &parseData)
	if err != nil {
		log.Error("Cannot parse response json")
		return res, err
	}

	if parseData.Status != "1" {
		log.WithFields(log.Fields{
			"info": parseData.Info,
		}).Error("Gaode return data error")
		return res, errors.New(parseData.Info)
	}

	if len(parseData.Geocodes) == 0 {
		log.Info("Map response is empty")
		return res, errors.New("Map response is empty")
	}

	geoCode := parseData.Geocodes[0].Location
	cityCode := parseData.Geocodes[0].Citycode
	address := parseData.Geocodes[0].FormattedAddress

	geoData := strings.Split(geoCode, ",")
	lng, err := strconv.ParseFloat(geoData[0], 64)
	if err != nil {
		log.Error("Cannot parse response geo location")
		return res, err
	}

	lat, err := strconv.ParseFloat(geoData[1], 64)
	if err != nil {
		log.Error("Cannot parse response geo location")
		return res, err
	}

	return MapLocation{
		Lng:      lng,
		Lat:      lat,
		CityCode: cityCode,
		Address:  address,
	}, nil
}

// MapLocation struct
type MapLocation struct {
	Lat      float64
	Lng      float64
	CityCode string
	Address  string
}

func (el MapLocation) String() string {
	return fmt.Sprintf("Addr: %v, Pos: %v, %v Code: %v", el.Address, el.Lng, el.Lat, el.CityCode)
}
