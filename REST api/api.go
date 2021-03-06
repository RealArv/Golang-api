package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

//Input struct for incoming json file
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

//Outgoing json result struct
type finalJsonStruct struct {
	TotalFlights        int    `json:"Total_Flights"`
	TotalMinutesDelayed int    `json:"Total_Minutes_Delayed"`
	TimeLabel           string `json:"Time_Label"`
}

//Requests
func main() {
	var inputData []inputJsonStruct
	output := []finalJsonStruct{}

	router := gin.Default()
	router.Use(cors.AllowAll())

	// GET request
	// router.GET("/get", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, output)
	// })

	router.POST("/post", func(c *gin.Context) {
		code := c.PostForm("code")
		file, _ := c.FormFile("file")

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		out, err := os.Open(file.Filename)
		byteValue, _ := ioutil.ReadFile(out.Name())
		errr := json.Unmarshal([]byte(byteValue), &inputData)

		for _, value := range inputData {
			if value.Airport.Code == code {
				output = append(output, finalJsonStruct{value.Statistics.Flights.Total, value.Statistics.MinutesDelayed.Total, value.Time.Label})
			}
		}

		// fmt.Println(output)

		if err != nil {
			fmt.Println(errr)
			c.JSON(500, gin.H{
				"error": "Some error is there",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"file": output,
		})

		output = nil
	})

	router.Run("localhost:9090")
}
