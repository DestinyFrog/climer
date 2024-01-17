package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Climer struct {
	Temperatura float32 `json:"temperatura"`
	Umidade		float32 `json:"umidade"`
}

func GetData() (data Climer, err error) {
	requestURL := "http://192.168.1.175/"

	res, err := http.Get(requestURL)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf( "http status error: %d", res.StatusCode )
		return
	}

	reqBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	json.Unmarshal( reqBody, &data )
	return
}