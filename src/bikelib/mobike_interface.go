package bikelib

import (
	"encoding/json"
	"errors"
	"net/url"

	log "github.com/sirupsen/logrus"
)

const mobikeURL = "https://mwx.mobike.com/mobike-api/rent/nearbyBikesInfo.do"

// Mobike struct
type Mobike struct {
	Lat float64
	Lng float64
}

// GetNearbyCar interface
func (mobike *Mobike) GetNearbyCar() ([]MobikeCar, error) {

	params := url.Values{}
	params.Add("latitude", "40.02015250763075")
	params.Add("longitude", "116.42243937431424")
	// params.Add("citycode", "010")

	nearbyData, err := PostFormData(mobikeURL, params)
	if err != nil {
		return nil, err
	}

	return mobike.parseJSON(nearbyData)
}

// parseJSON interface
func (mobike *Mobike) parseJSON(jsonData []byte) ([]MobikeCar, error) {
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

	return parseData.Object, nil
}
