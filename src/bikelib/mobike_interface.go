package bikelib

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const mobikeURL = "https://mwx.mobike.com/mobike-api/rent/nearbyBikesInfo.do"

// Mobike struct
type Mobike struct {
	Lng      float64
	Lat      float64
	CityCode string
}

// GetNearbyCar interface
func (mobike Mobike) GetNearbyCar() ([]BikeData, error) {

	params := url.Values{}
	lat := strconv.FormatFloat(mobike.Lat, 'f', 6, 64)
	lng := strconv.FormatFloat(mobike.Lng, 'f', 6, 64)
	params.Add("latitude", lat)
	params.Add("longitude", lng)
	params.Add("citycode", mobike.CityCode)

	nearbyData, err := PostFormData(mobikeURL, params)
	if err != nil {
		return nil, err
	}

	return mobike.parseJSON(nearbyData)
}

// parseJSON interface
func (mobike Mobike) parseJSON(jsonData []byte) ([]BikeData, error) {
	var parseData MobikeJSON

	err := json.Unmarshal(jsonData, &parseData)
	if err != nil {
		log.Error("Cannot parse response json")
		return nil, err
	}

	if parseData.Code != 0 {
		log.WithFields(log.Fields{
			"msg": parseData.Message,
		}).Error("Mobike return data error")
		return nil, errors.New(parseData.Message)
	}

	bikeRes := make([]BikeData, len(parseData.Object))
	for index, el := range parseData.Object {
		bikeRes[index] = BikeData{
			Lng:     el.DistX,
			Lat:     el.DistY,
			CarNo:   strconv.Itoa(el.DistNum),
			CarType: "mobike",
		}
	}

	return bikeRes, nil
}
