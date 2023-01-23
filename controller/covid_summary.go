package controller

import (
	"encoding/json"
	"io/ioutil"
	"lineman_asm_1/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CovidSummaryController(c *gin.Context) {
	// make HTTP request to URL
    resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

	// read response body into a byte slice
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// unmarshal response body into a Go struct
    var data models.DataStruct
    json.Unmarshal(body, &data)

	// setup response
	ret := models.Response {
		Province: make(map[string]int),
		AgeGroup: make(map[string]int),
	}
	// ret.AgeGroup["0-30"] = 0
	// ret.AgeGroup["31-60"] = 0
	// ret.AgeGroup["61+"] = 0
	// ret.AgeGroup["N/A"] = 0

	for _, d := range data.Data {
		if d.ProvinceEn != "" {
			if ret.Province[d.ProvinceEn] == 0 {
				ret.Province[d.ProvinceEn] = 1
			} else {
				ret.Province[d.ProvinceEn] += 1
			}
		} else {
			if _, ok := ret.Province["N/A"]; !ok {
				ret.Province["N/A"] = 1
			} else {
				ret.Province["N/A"] += 1 
			}
		}

		if string(d.Age) == "null" {
			ret.AgeGroup["N/A"] += 1
		} else {
			// convert byte array -> string
			// convert string -> int
			age, _ := strconv.Atoi(string(d.Age))
			
			if age >=0 && age <= 30 {
				ret.AgeGroup["0-30"] += 1
			} else if age >=31 && age <= 60 {
				ret.AgeGroup["31-60"] += 1
			} else if age >=61 {
				ret.AgeGroup["61+"] += 1
			}
		}
	}
	c.JSON(http.StatusOK, ret)
}
