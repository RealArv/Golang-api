package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type inputJsonStruct struct {
	Airport struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Airport"`
	Time struct {
		Label     string `json:"Label"`
		Month     int    `json:"Month"`
		MonthName string `json:"Month Name"`
		Year      int    `json:"Year"`
	} `json:"Time"`
	Statistics struct {
		OfDelays struct {
			Carrier                int `json:"Carrier"`
			LateAircraft           int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security               int `json:"Security"`
			Weather                int `json:"Weather"`
		} `json:"# of Delays"`
		Carriers struct {
			Names string `json:"Names"`
			Total int    `json:"Total"`
		} `json:"Carriers"`
		Flights struct {
			Cancelled int `json:"Cancelled"`
			Delayed   int `json:"Delayed"`
			Diverted  int `json:"Diverted"`
			OnTime    int `json:"On Time"`
			Total     int `json:"Total"`
		} `json:"Flights"`
		MinutesDelayed struct {
			Carrier                int `json:"Carrier"`
			LateAircraft           int `json:"Late Aircraft"`
			NationalAviationSystem int `json:"National Aviation System"`
			Security               int `json:"Security"`
			Total                  int `json:"Total"`
			Weather                int `json:"Weather"`
		} `json:"Minutes Delayed"`
	} `json:"Statistics"`
}

type finalJsonStruct struct {
	TotalFlights        int    `json:"Total Flights"`
	TotalMinutesDelayed int    `json:"Total Minutes Delayed"`
	TimeLabel           string `json:"Time Label"`
}

func main() {

	byteValue, _ := ioutil.ReadFile("airlines.json")
	var inputData []inputJsonStruct
	output := []finalJsonStruct{}
	err := json.Unmarshal([]byte(byteValue), &inputData)

	if err != nil {
		panic(err)
	}

	var temp string = "ATL"

	for _, value := range inputData {
		if value.Airport.Code == temp {
			output = append(output, finalJsonStruct{value.Statistics.Flights.Total, value.Statistics.MinutesDelayed.Total, value.Time.Label})
		}
	}

	finalJson, err := json.MarshalIndent(output, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
