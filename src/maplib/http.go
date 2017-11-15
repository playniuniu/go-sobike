package maplib

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// HTTPGet data
func HTTPGet(url string, params url.Values) ([]byte, error) {
	httpClient := &http.Client{}
	urlEncode := params.Encode()
	resp, err := httpClient.Get(apiURL + "?" + urlEncode)

	if err != nil {
		log.WithFields(log.Fields{
			"url": url,
		}).Error("Cannot get data")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"code": resp.StatusCode,
		}).Error("Status code is")
		return nil, errors.New("Status code error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Parse response error")
		return nil, errors.New("Parse response error")
	}

	return body, nil
}
