package bikelib

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const ofoURL = "https://san.ofo.so/ofo/Api/nearbyofoCar"
const ofoToken = "0ABE7990-A5A9-11E6-8FD5-016BD2CF67D2"

// Ofobike struct
type Ofobike struct {
	Lat float64
	Lng float64
}

// GetNearbyCar get all nearby car
func (ofo *Ofobike) GetNearbyCar() ([]OfoCar, error) {

	params := url.Values{}
	lat := strconv.FormatFloat(ofo.Lat, 'f', 6, 64)
	lng := strconv.FormatFloat(ofo.Lng, 'f', 6, 64)
	params.Add("lat", lat)
	params.Add("lng", lng)
	params.Add("token", ofoToken)
	params.Add("source", "0")

	nearbyData, err := PostFormData(ofoURL, params)
	if err != nil {
		return nil, err
	}

	return ofo.parseJSON(nearbyData)
}

func (ofo *Ofobike) parseJSON(jsonData []byte) ([]OfoCar, error) {
	var parseData OfoJSON

	err := json.Unmarshal(jsonData, &parseData)
	if err != nil {
		log.Error("Cannot parse response json")
		return nil, err
	}

	if parseData.ErrorCode != 200 {
		log.WithFields(log.Fields{
			"msg": parseData.Msg,
		}).Error("Ofo return data error")
		return nil, errors.New(parseData.Msg)
	}

	return parseData.Values.Info.Cars, nil
}
