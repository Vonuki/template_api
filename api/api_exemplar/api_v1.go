package api_exemplar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiUrl = "http://worldtimeapi.org/api/timezone/Europe/%s"

type StructResponse struct {
	Time string `json:"datetime"`
}

type ExamplarTimeAPI struct {
	comment string
}

func NewBestTimeAPI() *ExamplarTimeAPI {
	return &ExamplarTimeAPI{comment: "test"}
}

// handle request to API
func (*ExamplarTimeAPI) GetTime(city string) (string, error) {
	fmt.Printf("Request to URL %s \n", apiUrl)

	// request to API
	resp, err := http.Get(fmt.Sprintf(apiUrl, city))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// get body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// load to struct
	structResponse := StructResponse{}
	err = json.Unmarshal(body, &structResponse)
	if err != nil {
		return "", err
	}

	return structResponse.Time, nil
}
