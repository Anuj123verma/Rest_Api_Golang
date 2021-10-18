package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Anuj123Verma/Rest_Api_Golang/entity"
)

func ConsumeApi(states entity.CovidCases) entity.CovidCases {
	response, err := http.Get("https://api.rootnet.in/covid19-in/unofficial/covid19india.org/statewise")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &states)
	return states
}
