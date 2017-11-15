package maplib

import (
	"encoding/json"
	"errors"
	"net/url"
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

// Bike struct
type Bike struct {
	Lat string
	Lng string
}

// GetGeoLoc function
func (mapaddr *MapAddr) GetGeoLoc() (Bike, error) {
	var res Bike

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

func (mapaddr *MapAddr) parseJSON(jsonData []byte) (Bike, error) {
	var res Bike
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
		log.Error("Map addr is empty")
		return res, errors.New("Map addr is empty")
	}

	geoCode := parseData.Geocodes[0].Location
	geoData := strings.Split(geoCode, ",")

	return Bike{Lng: geoData[0], Lat: geoData[1]}, nil
}
