// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// type airportStruct struct {
// 	Airport struct {
// 		Code string `json:"Code"`
// 		Name string `json:"Name"`
// 	} `json:"Airport"`
// 	Time struct {
// 		Label     string `json:"Label"`
// 		Month     int    `json:"Month"`
// 		MonthName string `json:"Month Name"`
// 		Year      int    `json:"Year"`
// 	} `json:"Time"`
// 	Statistics struct {
// 		OfDelays struct {
// 			Carrier                int `json:"Carrier"`
// 			LateAircraft           int `json:"Late Aircraft"`
// 			NationalAviationSystem int `json:"National Aviation System"`
// 			Security               int `json:"Security"`
// 			Weather                int `json:"Weather"`
// 		} `json:"# of Delays"`
// 		Carriers struct {
// 			Names string `json:"Names"`
// 			Total int    `json:"Total"`
// 		} `json:"Carriers"`
// 		Flights struct {
// 			Cancelled int `json:"Cancelled"`
// 			Delayed   int `json:"Delayed"`
// 			Diverted  int `json:"Diverted"`
// 			OnTime    int `json:"On Time"`
// 			Total     int `json:"Total"`
// 		} `json:"Flights"`
// 		MinutesDelayed struct {
// 			Carrier                int `json:"Carrier"`
// 			LateAircraft           int `json:"Late Aircraft"`
// 			NationalAviationSystem int `json:"National Aviation System"`
// 			Security               int `json:"Security"`
// 			Total                  int `json:"Total"`
// 			Weather                int `json:"Weather"`
// 		} `json:"Minutes Delayed"`
// 	} `json:"Statistics"`
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/airport")
// 	router.Run("localhost:9090")
// }
