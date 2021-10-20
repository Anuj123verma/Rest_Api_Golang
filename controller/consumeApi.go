package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Anuj123Verma/Rest_Api_Golang/entity"
)

func ConsumeApi(states entity.CovidCases) entity.CovidCases {
	// getting the response from the api
	response, err := http.Get("https://api.rootnet.in/covid19-in/unofficial/covid19india.org/statewise")
	// error then exit
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// getting the responseData
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Unmarshalling the response data
	json.Unmarshal(responseData, &states)

	// returning the data stored in CovidCases struct
	return states
}
