package entity

import (
	"time"
)

// CovidCases entity declared for storing the data received from consumed api
type CovidCases struct {
	Success bool `json:"success"`
	Data    struct {
		Source        string    `json:"source"`
		LastRefreshed time.Time `json:"lastRefreshed"`
		Total         struct {
			Confirmed int `json:"confirmed"`
			Recovered int `json:"recovered"`
			Deaths    int `json:"deaths"`
			Active    int `json:"active"`
		} `json:"total"`
		Statewise []struct {
			State     string `json:"state" bson:"state"`
			Confirmed int    `json:"confirmed" bson:"confirmed cases"`
			Recovered int    `json:"recovered" bson:"recovered cases"`
			Deaths    int    `json:"deaths" bson:"deaths"`
			Active    int    `json:"active" bson:"active cases"`
		} `json:"statewise"`
	} `json:"data"`
	LastRefreshed    time.Time `json:"lastRefreshed"`
	LastOriginUpdate time.Time `json:"lastOriginUpdate"`
}
